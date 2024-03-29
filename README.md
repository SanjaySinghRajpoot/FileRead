
# FileRead Web Application
FileRead is a web application built with Go (Golang) using the Gin web framework. It provides functionalities to serve files, display a welcome message. This README provides an overview of the project structure and instructions for setup and usage.


## Installation


1. Run the following command from the root directory to start the Docker container:

```bash
docker-compose up
```


2. Running Directly with Go
If you prefer to run the application directly with Go, follow these steps:

    1. Ensure you have Go installed on your system. If not, you can download and install it from the official Go website: Install Go

    2. Navigate to the root directory of your project.

    3. Run the following command to start the application:

    ```bash 
    go run main.go
    ```

    4. Your application will be accessible at http://localhost:8080

3. After completing the installation, running the `/data?n={file_name}&m={line_number}` endpoint requires adding a text file in the `tmp/data` directory. The file should be named in the format `{file_name}.txt`, where `{file_name}` represents the file number. A sample `1.txt` file is added in the repo, please add more file in the same folder for testing. 

- For example, if you want to test with file number 1, the file should be named `1.txt`. Ensure the text file contains content to read, and then make a request to the `/data` endpoint with the appropriate query parameters (`n` for the file name and `m` for the line number). This will allow you to verify the functionality of the endpoint by retrieving the specified line from the text file.

## Testing

To run tests, execute `go test` from the root directory of the project. Currently, the `TestGetDataForLine` test will fail due to the need for replacing the `mockResponse` with the actual content present on the specified line using the `m` parameter.

### Docker Compose testing
The Docker image was tested using the Docker Compose configuration with resource constraints set to **1.5 GB** of RAM and **1 CPU core**. Below is the screenshot of the test result:

The test was successful, demonstrating that the Docker image runs within the specified resource constraints without any issues. This ensures that the application can be deployed in environments with limited resources while maintaining stability and performance.
![Screenshot from 2024-01-26 16-12-19](https://github.com/SanjaySinghRajpoot/FileRead/assets/67458417/2cd6e1b1-e99b-4bf6-81b1-4d144f812f7e)


## File Streaming Optimization

#### Purpose:
The optimization aims to enhance the efficiency and simplicity of streaming entire files in a web application developed using the Gin web framework in Go.

#### Method:
Utilizes the `ctx.File(filePath)` method provided by the Gin framework, which internally leverages the `http.ServeFile` function from the Go standard library.

#### Why Use This Optimization:
1. **Efficiency**: The `http.ServeFile` function efficiently serves file contents to clients by utilizing system-level optimizations, reducing resource usage and improving performance.
2. **Simplicity**: By using `ctx.File(filePath)`, developers can achieve file streaming with minimal code, enhancing code readability and maintainability.
3. **Security**: The implementation provided by `http.ServeFile` handles various security considerations, such as preventing directory traversal attacks, ensuring safe file serving.
4. **Consistency**: Using standard library functions promotes consistency across different parts of the codebase and aligns with best practices recommended by the Go community.
5. **Robustness**: Leveraging well-tested standard library functionality reduces the likelihood of bugs and enhances the reliability of the application.

#### Conclusion:
By employing the `ctx.File(filePath)` method, the optimization simplifies the code while ensuring efficient and secure file streaming, contributing to a more robust and maintainable web application.
