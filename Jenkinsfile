pipeline {
    agent any

    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'main', url: 'https://github.com/AthulKrishna/ci-cd.git'
            }
        }

        stage('Install Dependencies') {
            steps {
                sh 'go mod tidy'
            }
        }

        stage('Build App') {
            steps {
                sh 'GOOS=linux GOARCH=amd64 go build -o app .'
            }
        }

        stage('Run App') {
            steps {
                sh './app &'
                sleep 5
            }
        }

        stage('Test the App') {
            steps {
                script {
                    sh 'curl http://localhost:8080/health'
                }
            }
        }

        stage('Run Tests') {
            steps {
                sh 'go test -v ./...'
            }
        }

        stage('Cleanup') {
            steps {
                sh 'pkill -f "app"'
            }
        }
    }
}
