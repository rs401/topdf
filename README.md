<div id="top"></div>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/rs401/topdf">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">ToPDF</h3>

  <p align="center">
    ToPDF offers an API endpoint that attempts to convert any file to PDF.
    <br />
    <a href="https://github.com/rs401/topdf"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/rs401/topdf">View Demo</a>
    ·
    <a href="https://github.com/rs401/topdf/issues">Report Bug</a>
    ·
    <a href="https://github.com/rs401/topdf/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
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
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](https://example.com)

ToPDF offers an API endpoint that attempts to convert any file to PDF. A microservice in a docker container.


<p align="right">(<a href="#top">back to top</a>)</p>



### Built With

* [Go](https://go.dev/)
* [Docker](https://www.docker.com/)
* [Pandoc](https://pandoc.org/)
* [Alpine](https://www.alpinelinux.org/)
* [TinyTeX 👑](https://yihui.org/tinytex/)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

Build main, build the docker image or pull the image when I upload it.

### Prerequisites

Go(lang) to build a binary or Docker to build the docker image. If building the go binary, you will also need Pandoc installed with pdflatex.

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/rs401/topdf.git
   ```
2. Build binary
   ```sh
   go build .
   PORT=8888 ./topdf
   ```

OR


2. Build the Docker image
   ```sh
   docker build -t rs401/topdf:1.0 .
   docker run --rm -d -p 8888:8888 --name topdf rs401/topdf:1.0
   ```

   OR, if you would like to change the API port.

   ```sh
   docker run --rm -d -e PORT=9999 --expose 9999 -p 9999:9999 --name topdf rs401/topdf:1.0
   ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

```sh
curl --request POST http://127.0.0.1:8888/topdf -F file=@main.go --output out.pdf
```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Parameterize the listening port
- [x] Load config from env file

See the [open issues](https://github.com/rs401/topdf/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Richard Stadnick - [![LinkedIn][linkedin-shield]][linkedin-url]

Project Link: [https://github.com/rs401/topdf](https://github.com/rs401/topdf)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments


<p align="right">(<a href="#top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/rs401/topdf.svg?style=for-the-badge
[contributors-url]: https://github.com/rs401/topdf/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/rs401/topdf.svg?style=for-the-badge
[forks-url]: https://github.com/rs401/topdf/network/members
[stars-shield]: https://img.shields.io/github/stars/rs401/topdf.svg?style=for-the-badge
[stars-url]: https://github.com/rs401/topdf/stargazers
[issues-shield]: https://img.shields.io/github/issues/rs401/topdf.svg?style=for-the-badge
[issues-url]: https://github.com/rs401/topdf/issues
[license-shield]: https://img.shields.io/github/license/rs401/topdf.svg?style=for-the-badge
[license-url]: https://github.com/rs401/topdf/blob/main/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/richard-stadnick-3b4ab53b
[product-screenshot]: images/screenshot.png