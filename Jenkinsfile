pipeline {
    agent {label 'docker'} 
    // added sept 15t chg A
    // added sept 18th chg B
    // added chg one
    // add one to master
    // add two to master
    stages {
        stage('build and test') {
            agent { 
                docker { 
                    label 'docker'
                    image 'golang:1.14' 
			// added one to feature
			// added two to featutue
			// added three to feature
                        }
                             }
            environment {
                GOCACHE = '/tmp/gocache'
            }
            steps {
                sh 'go build'
                sh 'go test ./...'
                sh 'echo test1'
            }
        }
        stage('deploy') {
            agent any
            environment {
                BUILD_NUMBER_BASE = '0'
                VERSION_MAJOR = '0'
                VERSION_MINOR = '1'
                VERSION_PATCH = "${env.BUILD_NUMBER.toInteger() - BUILD_NUMBER_BASE.toInteger()}"
            }
            when {
                branch 'master'
            }
            steps {
                script {
                    docker.withRegistry('', 'dockerhub') {
                        def customImage = docker.build("nicholasgull/hello-jenkins:$VERSION_MAJOR.$VERSION_MINOR.$VERSION_PATCH")
                        customImage.push()
                    }
                }
            }
        }
    }
}
