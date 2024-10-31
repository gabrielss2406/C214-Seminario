pipeline {
    
    agent any

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/gabrielss2406/C214-Seminario'
            }
        }

        stage('Build') {
            steps {
                sh 'go build ./...'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./... -v'
            }
        }
    }

    post {
        always {
            echo 'Pipeline completo!'
        }
    }
}
