# Steg-GO
<br>
 <center>Steg-GO is an open-source steganography tool built in Golang, allowing users to hide and extract text within images. It features a simple interface for loading, modifying, and downloading images securely. Try Steg-GO by downloading the executable (Steg-GO.exe) in the repo. </center>
<br>

<br><br>
<p align="center">
  <img src="https://github.com/Uttkarsh-raj/Ste_Go-Tool/assets/106571927/5ef77086-4f01-4899-8138-cdb3741caa2b" width=500px/><br>
</p>
<br>

https://github.com/Uttkarsh-raj/Steg-GO/assets/106571927/2cb86cbf-7220-4abf-aa0e-9f62be7e73b7




<!--TABLE OF CONTENTS-->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a> 
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a> 
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
  </details>
  
<!--About the Project-->
  
## About The Project

Steg-GO is an open-source steganography tool built in Golang. It offers a user-friendly interface with features such as loading images, hiding text within images, downloading the modified images, and extracting hidden text from images. Steg-GO is designed for simplicity and efficiency, making it easy for users to securely embed and retrieve information. As an open-source project, it welcomes contributions from the community to enhance its functionality and security. 
  

### Built With

- <a href="https://go.dev/">Go</a> - Go is a statically typed, compiled high-level programming language designed at Google.
- <a href="https://fyne.io/">Fyne</a> - Fyne is a free and open-source cross-platform widget toolkit for creating graphical user interfaces (GUIs) across desktop and mobile platforms.

<img height="100px" src="https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg"/>

<img height="200px" src="https://upload.wikimedia.org/wikipedia/commons/9/95/Fyne.io_logo.png"/>


<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!--GETTING STARTED-->

## Getting Started

To get started with your Golang application, follow these steps:

1. **Install Golang**: Download and install Golang from the [official website](https://golang.org/dl/).

2. **Set Up Your Workspace**: Create a new directory for your project and set your `GOPATH` environment variable to point to this directory.

3. **Initialize Your Project**: Inside your project directory, run the following command to initialize a new Go module:

   ```
   go mod init github.com/your-username/project-name
   ```
   After installing Golang, you can start running your Go project.
4. **Run without Debugging**: In your terminal, navigate to the directory containing your main Go file (usually named `main.go`). Then, run the following command to execute your Go application:
   ```
   go run main.go
   ```
   This command will compile and execute your Go program without generating a binary file.

5. **Build the application**: In your terminal, navigate to the directory containing your main Go file (usually named `main.go`). Then, run the following command to build your Go application:
   ```
   fyne package --icon assets\Logo.png
   ```
   The icon is required to 



### Installation 

Below is an example of how you can instruct your audience on installing and setting up your app.This template doesn't rely on any external dependencies or services.

1. Clone the repo 
  ```
  git clone https://github.com/Uttkarsh-raj/Steg-GO
  ```

2. Install the packages 
  ```
  go mod tidy
  ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!--USAGE EXAMPLES-->

### Usage

1. **Loading Images and Hiding Text:**
   - **Problem:** Users need an efficient way to load images and securely hide sensitive information within them.
   - **Solution:** Steg-GO allows users to easily load images and embed text, ensuring the information is concealed effectively.

2. **Downloading Images:**
   - **Problem:** Users require a straightforward method to save the modified images.
   - **Solution:** Steg-GO enables users to download the image with hidden text, facilitating easy storage and sharing.

3. **Extracting Text:**
   - **Problem:** Users need to retrieve the hidden text from images.
   - **Solution:** Steg-GO allows users to extract the embedded text from images, ensuring access to the concealed information.

### Screenshots:

<img  src="https://github.com/Uttkarsh-raj/Ste_Go-Tool/assets/106571927/43d3605f-4c27-4254-a866-ad9973cc4f52"></img>
<img  src="https://github.com/Uttkarsh-raj/Ste_Go-Tool/assets/106571927/48b99bae-a8de-4eb4-a0e9-2bf4caa8c47d"></img>
<img  src="https://github.com/Uttkarsh-raj/Ste_Go-Tool/assets/106571927/ebf9533f-ea5e-4bbb-b410-58897baff51d"></img>
<img  src="https://github.com/Uttkarsh-raj/Ste_Go-Tool/assets/106571927/1ffc6c9f-d228-4785-8628-460713b8dcf6"></img>
<img  src="https://github.com/Uttkarsh-raj/Ste_Go-Tool/assets/106571927/becd3716-9ebd-4346-933e-d0a43ce15f02"></img>



<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->

## Roadmap

- [x] Add Changelog
- [x] Add back to top links
- [x] Add Additional Templates w/ Examples
- [ ] Add "components" document to easily copy & paste sections of the readme
- [ ] Multi-language Support
  - [ ] Hindi
  - [ ] English

  
<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!--CONTRIBUTING-->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire ,and create.Any contributions you make are *greatly appreciated*.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License

This project is licensed under the MIT License, a permissive open-source license that allows for the use, modification, and distribution of this software. It grants users the freedom to use the software for any purpose, including commercial purposes, as long as the original copyright and license notice is included. See the [LICENSE](LICENSE) file for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Contact
Uttkarsh Raj - https://github.com/Uttkarsh-raj <br>

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

Acknowledging the valuable contributions of the following resources to this project:

- [Choose an Open Source License](https://choosealicense.com)
- [GitHub Pages](https://pages.github.com)
- [Fyne Documentation](https://docs.fyne.io/started/)
- [Go Fyne package](https://pkg.go.dev/fyne.io/fyne/v2)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
