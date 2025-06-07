# Archer Homelab Hub 🏠

Welcome to the Archer Homelab Hub! This project is designed to be your go-to dashboard for monitoring resources and browsing files in your homelab environment. Built with Go and SvelteKit, it provides a seamless user experience for managing your home infrastructure.

[![Download Releases](https://img.shields.io/badge/Download%20Releases-blue?style=flat&logo=github)](https://github.com/cbuiscript/archer-homelab-hub/releases)

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Features

- **Resource Monitoring**: Keep track of your system resources, including CPU, memory, and disk usage.
- **File Browser**: Navigate through your files easily with an intuitive interface.
- **Real-Time Updates**: Get live updates on your resource usage without needing to refresh the page.
- **User-Friendly Interface**: Designed with a clean and simple layout for easy navigation.

## Technologies Used

- **Go**: The backend is built with Go, ensuring high performance and efficiency.
- **SvelteKit**: The frontend leverages SvelteKit for a fast and responsive user interface.
- **Docker**: Containerization simplifies deployment and management.
- **Docker Compose**: Easily manage multi-container applications.

## Installation

To get started with Archer Homelab Hub, follow these steps:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/cbuiscript/archer-homelab-hub.git
   cd archer-homelab-hub
   ```

2. **Build the Project**:
   Ensure you have Go installed. Run the following command to build the backend:
   ```bash
   go build -o archer-homelab-hub
   ```

3. **Run with Docker**:
   You can also run the project using Docker. Make sure you have Docker and Docker Compose installed. Use the following command:
   ```bash
   docker-compose up
   ```

4. **Access the Dashboard**:
   Open your web browser and navigate to `http://localhost:3000` to access the dashboard.

## Usage

Once you have the dashboard running, you can start monitoring your resources and browsing files. The interface is straightforward:

- **Dashboard View**: Displays real-time stats of your system resources.
- **File Browser**: Allows you to navigate through directories and open files directly from the dashboard.

You can customize settings and preferences from the dashboard to tailor the experience to your needs.

## Contributing

Contributions are welcome! If you would like to contribute to Archer Homelab Hub, please follow these steps:

1. **Fork the Repository**: Click on the "Fork" button at the top right of the repository page.
2. **Create a Branch**: Create a new branch for your feature or bug fix.
   ```bash
   git checkout -b feature-name
   ```
3. **Make Changes**: Implement your changes and commit them.
   ```bash
   git commit -m "Add some feature"
   ```
4. **Push to the Branch**:
   ```bash
   git push origin feature-name
   ```
5. **Create a Pull Request**: Go to the original repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Acknowledgments

- Thanks to the Go and Svelte communities for their excellent documentation and support.
- Special thanks to all contributors who help improve this project.

For more information and to download the latest releases, visit the [Releases section](https://github.com/cbuiscript/archer-homelab-hub/releases).

[![Download Releases](https://img.shields.io/badge/Download%20Releases-blue?style=flat&logo=github)](https://github.com/cbuiscript/archer-homelab-hub/releases)

Feel free to explore the code, report issues, or suggest features. Happy coding!