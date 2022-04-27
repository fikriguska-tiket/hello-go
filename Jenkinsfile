node ('master') {   
    stage('Hello') {
        sh 'echo "this is should be a build command"'
        sh 'hostname && pwd && whoami'
    }
    stage('gcloud') {
        sh 'gcloud'
    }   
}
