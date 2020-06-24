# Authedia Developers

The Authedia Protocol sets the standard for proving the authenticity of digital media. This service is freely available to developers. The Authedia Protocol can be leveraged for [proof of ownership](https://github.com/Authedia/Developers#wrap-media) - registering media originated from your application. As well as [verifying the authenticity](https://github.com/Authedia/Developers#verify-media) of digital media.


### Getting Started
1. Register for an API key
  * Create an account and register for a key [here](www.authedia.com)
2. Call the Authedia API in your application

### Supported File Types
* Audio
  * wav
* Images
  * png
  * jpeg
* Video
  * mp4

### Provided Languages
* [Python](https://github.com/Authedia/Developers/tree/python)
* [Go](https://github.com/Authedia/Developers/tree/go)

# Documentation

### Wrap Media
* Inserts secret information into the media file which can later be unwrapped to prove ownership
  * Note: Wrapping media with the API will not prove authenticity, only ownership. As we are not able to ascertain the origin of the media file. To capture verifiable authentic media use our [mobile app](www.authedia.com)

```
import authedia

authedia.wrap(
    api_key=your_api_key,
    input_file_name='example_media/PNG.png',
    output_file_name='test.png'
)
```

**wrap(api_key, input_file_name, output_file_name)**  
  * api_key : (str) : User's API key
  * input_file_name : (str) : Path to input file
  * output_file_name : (str) : Desired location of output file

Returns  
* wrap_response : (dict) : See below for key value info

```
{
    'Success' : bool : successfully processed
    'Message' : str  : error message if request is unsuccessful
}
```

### Verify Media
* Unwraps secret information from the media file. The ownership and authenticity will be determined. Detailed information regarding the originality of the media will be returned.

```
response = authedia.verify(
    api_key=your_api_key,
    input_file_name='test.png'
)
```

**verify(api_key, input_file_name)**  
  * api_key : (str) : User's API key
  * input_file_name : (str) : Path to input file

Returns  
  * verify_response : (dict) : See below for key value info
  ```
  {
      'Bounds'  : str  : Original size of media
      'Success' : bool : Successfully processed
      'Message' : str  : Error message if request is unsuccessful
      'VerifiedBlocks' : list[ [bool, x, y] ] : successfully verified. (x, y) position in the data
      'DataBlock' : dict : {
          'DateTime' : str  : Date time media was captured
          'Device'   : str  : Type of device media was captured on
          'Location' : dict : {
              'LatLong' : str : 'lat,long'
              'Time' : str : Time the location was captured
          }
      }
  }
  ```
