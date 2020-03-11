node{
    git credentialsId: 'f92c1ee0-9d73-406a-bf55-b8c8bc6e6cd3', url: 'https://github.com/Mingo14/Capstone.git'
    docker.withRegistry('https://registry.hub.docker.com/', 'docker2'){
        def dockerfile = 'Dockerfile'
        def image = docker.build("mbradfield/capstone-go-app:1.${env.BUILD_ID}")
        image.push()
    }
    sh "docker tag 25c4671a1478 gcr.io/decoded-doodad-265918capstone-go-app:1.${env.BUILD_ID}"
    sh "gcloud auth configure-docker"
    sh "docker-credential-gcr configure-docker"
    sh "gcloud auth activate-service-account jenkins@decoded-doodad-265918.iam.gserviceaccount.com --key-file=jenkins-sa.json"
    sh "docker login -u oauth2accesstoken -p \"\$(gcloud auth print-access-token)\" https://gcr.io"
    sh "docker push gcr.io/decoded-doodad-265918/capstone-go-app:1.${env.BUILD_ID}"
}