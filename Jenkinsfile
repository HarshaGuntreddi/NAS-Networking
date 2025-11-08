pipeline {
  agent any

  options {
    timestamps()
    ansiColor('xterm')
    buildDiscarder(logRotator(numToKeepStr: '15'))
    timeout(time: 15, unit: 'MINUTES')
  }

  stages {
    stage('Checkout') {
      steps {
        deleteDir()
        checkout scm
      }
    }

    stage('Go Build') {
      steps {
        sh 'go version || true'
        sh 'if command -v go >/dev/null 2>&1; then go build nas_portal.go; fi'
      }
    }

    stage('Ansible Lint Stub') {
      steps {
        sh 'echo "run ansible-lint site.yml here"'
      }
    }

    stage('Docker Build Stub') {
      steps {
        sh 'echo "docker build -f Dockerfile.webapp ."'
      }
    }
  }

  post {
    always {
      archiveArtifacts artifacts: 'nas_portal README.md *.yml *.md', fingerprint: true
    }
  }
}
