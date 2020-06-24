# Authedia Developers
Using the Authedia API in Python. You can find the documentation [here](https://github.com/Authedia/Developers#documentation)

### Example
```
import authedia

authedia.Wrap(
    api_key=your_api_key,
    input_file_name='example_media/PNG.png',
    output_file_name='test.png'
)

response = authedia.Verify(
    api_key=your_api_key,
    input_file_name='test.png'
)
```
