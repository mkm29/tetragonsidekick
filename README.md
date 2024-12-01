# tetragonsidekick

The Tetragonsidekick project is a Kubernetes operator designed to process logs emitted by Tetragon pods, enhance the logs using an LLM (Large Language Model), and forward them to an OpenSearch server for future analysis. This project uses Redpands as the central data platform for event streaming, ensuring scalability and real-time processing.

## About The Project

### Key Features

- **Log Streaming**: Captures logs from Tetragon pods in the `kube-system` namespace and streams them into a Redpanda topic.
- **Protobuf Serialization**: Processes Tetragon logs using the official `events.proto` definitions.
- **Event Enhancement**: Uses an LLM (e.g., OpenAI GPT-4) to analyze and enhance events.
- **OpenSearch Integration**: Forward enhanced logs to OpenSearch for indexing and analytics.
- **Scalable Architecture**: Built on Redpanda for efficient and scalable data processing.

### Architecture Overview

1. **Tetragon Pods**:
   1. Emit structured logs in protobuf format.
2. **Custom Kubernetes Controller**:
   1. Captures logs from Tetragon pods.
   2. Serializes logs using Tetragon’s protobuf definitions.
   3. Streams logs to Redpands.
3. **Stream Processor**:
   1. Enhances logs using an LLM.
   2. Writes enhanced logs back to a separate Redpands topic.
4. **OpenSearch Sink**:
   1. Consumes enhanced logs from Redpands.
   2. Indexes them in OpenSearch.

#### Event Structure

The official Tetragon [protobuf](https://github.com/cilium/tetragon/blob/main/api/v1/tetragon/events.proto) for events provides a structured way to work with Tetragon logs. By leveraging this, you can deserialize and process events effectively, ensuring compatibility with Tetragon's data structure. Below is a detailed implementation plan using the protobuf definitions and integrating Redpands for a data-driven architecture. Sample [event.json](./assets/event.json).

## Getting Started

### Prerequisites

1. Kubernetes Cluster: A running Kubernetes cluster (v1.29 or higher).
   1. Linux kernel >= 4.19 (for eBPF)
2. Redpands: Installed in the cluster.
3. OpenSearch: Running in the cluster or externally.
4. Operator SDK: Installed on your development machine.
5. Protobuf Compiler: `protoc` must be installed.

### Installation

_todo_

### Usage

_todo_

## Roadmap

_todo_

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

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Mitchell Murphy - [@mitchellmurphy](https://www.linkedin.com/in/mitchellmurphy/) | [mitchmurphy.io](https://mitchmurphy.io/)

Project Link: [https://github.com/mkm29/tetragonsidekick](https://github.com/mkm29/tetragonsidekick)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

* [Choose an Open Source License](https://choosealicense.com)
* [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
* [Malven's Flexbox Cheatsheet](https://flexbox.malven.co/)
* [Malven's Grid Cheatsheet](https://grid.malven.co/)
* [Img Shields](https://shields.io)
* [GitHub Pages](https://pages.github.com)
* [Font Awesome](https://fontawesome.com)
* [React Icons](https://react-icons.github.io/react-icons/search)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
