import os
import json
import base64
import requests
import tempfile

URL = "https://authedia.uc.r.appspot.com/"

RESPONSE_DICT = {
    'Success' : False,
    'Message' : 'Bad response'
}

def wrap(api_key, input_file_name, output_file_name=''):
    '''
    PARAMS
        api_key : (str) : User's API key
        input_file_name : (str) : Path to input file
        output_file_name : (str) : Desired location of output file

    RETRUNS
        wrap_response : (dict) : See documentation for details
    '''

    if api_key == '':
        return 'API key required', False

    if output_file_name == '':
        dirs = input_file_name.split('/')

        if len(dirs) > 1:

            output_file_name = os.path.join(
                dirs[:-1].join('/'),
                'authedia_' + input_file_name
            )

        else:

            output_file_name = 'authedia_' + input_file_name

    # LOAD FILE
    with open(input_file_name, 'rb') as file:

        response = requests.post(
            URL + 'wrap_request',
            files=dict(data=file),
            data={'api_key' : json.dumps({
                'ApiKey': api_key,
            })}
        )

    try:
        response_json = response.json()

        # CHECK FOR SUCCESSFUL RESPONSE
        if response_json.get('Success', False):

            # GET UTF-8 STRING FROM RESPONSE
            byte_string = response_json['Bytes'].split(',')
            byte_string = byte_string[-1]

            # UTF-8 TO BASE64 STRING
            file_base64 = base64.b64decode(byte_string.decode('utf-8'))

            # WRITING BASE64 STRING TO FILE
            with open(output_file_name, 'wb') as f:
                f.write(file_base64)

        response_json.pop('Bytes', None)

        return response_json

    except:
        return RESPONSE_DICT


def verify(api_key, input_file_name):
    '''
    PARAMS
        api_key : (str) : User's API key
        input_file_name : (str) : Path to input file

    RETRUNS
        verify_request : (dict) : See documentation for details
    '''
    # LOAD
    with open(input_file_name, 'rb') as file:

        response = requests.post(
            URL + 'verify_request',
            files=dict(data=file),
            data={'api_key' : json.dumps({
                'ApiKey': api_key,
            })}
        )

    try:
        return response.json()

    except:
        return RESPONSE_DICT
