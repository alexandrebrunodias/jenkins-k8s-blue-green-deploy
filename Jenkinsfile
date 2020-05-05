pipeline {
	agent any
	
	tools {
		go 'go-1.14'
	}

	environment {
		GO111MODULE = 'on'
		tag = "alexdias/store:$BUILD_NUMBER"
	}
	
	stages {
		stage('Install dependencies') {
			steps {
					sh 'go mod download'
			}
		}
		stage('Lint') {
			steps {
					sh 'golint ./...'
			}
		}
		stage('Test') {
			steps {
					sh 'go test -v ./...'
			}
		}

		stage('Docker build, push') {
			steps{
				script{
					image = docker.build tag
					image.push()
					image.push("latest")
				}

			}
		}

		// stage('Deploy'){
		// 	steps{
		// 		withAWS(credentials: 'aws', region: 'us-west-1')
		// 		{
		// 			sh 'aws eks --region=us-west-1 update-kubeconfig --name capstone'
		// 			sh 'kubectl apply -f eks-deployment.yml'
		// 		}
		// 	}
		// }
	}
}
