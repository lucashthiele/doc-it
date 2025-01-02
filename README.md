# 🚀 DocIT

DocIT is a powerful tool designed to simplify the process of creating professional and comprehensive documentation for your GitHub repositories. With the help of an AI-powered agent, DocIT automates the process of analyzing your codebase and generating a detailed README file that highlights all essential information about your project.  

## ✨ Features  

- **Seamless GitHub Integration**: Connect your GitHub account and choose the repository to document.  
- **AI-Powered Analysis**: The AI agent examines your repository, extracting key details about your project.  
- **Detailed Documentation**: Automatically generate a README file with sections like:  
  - Project title and brief description  
  - Installation instructions  
  - Build and run commands  
  - Contribution guidelines  
  - Any additional information the AI deems relevant  
- **Customizable Output**: Edit the generated documentation to suit your needs before automagically adding it to your repository.  

## 🛠️ Getting Started  

### Prerequisites  
- Go 1.23 or later installed on your system  
- A GitHub account with access to the repositories you want to document  

### Installation  

1. Clone the repository:  
   ```bash  
   git clone https://github.com/lucashthiele/doc-it.git  
   cd doc-it  
   ```  

2. Install dependencies:  
   ```bash  
   go mod tidy  
   ```  

### Usage  

1. Run the application:  
   ```bash  
   make run
   ```  

2. Follow the instructions in the terminal to authenticate your GitHub account.  

3. Select a repository, and let the AI agent do its magic.  

4. Customize the generated README if needed and publish it directly to your repository.  

### Build the App  

To build the application:  
```bash  
go build -o doc-it  
```  

Run the built binary:  
```bash  
./doc-it  
```  

## 🤝 Contributing  

Contributions are welcome! If you’d like to improve DocIT, please:  
1. Fork the repository  
2. Create a new branch (`git checkout -b feature/amazing-feature`)  
3. Commit your changes (`git commit -m 'Add some amazing feature'`)  
4. Push to the branch (`git push origin feature/amazing-feature`)  
5. Open a Pull Request  

## 📜 License  

This project is licensed under the MIT License. See the [LICENSE](https://github.com/lucashthiele/doc-it/blob/main/LICENCE) file for details.  

---

Unleash the potential of your repositories with **DocIT**—because every great project deserves great documentation!  
