FROM gradle as build

ENV APP_HOME=/apps

WORKDIR $APP_HOME

COPY build.gradle settings.gradle gradlew $APP_HOME

COPY gradle $APP_HOME/gradle

RUN chmod +x gradlew

RUN ./gradlew build || return 0

COPY src $APP_HOME/src

RUN ./gradlew clean build

FROM openjdk

ENV APP_HOME=/apps
ARG ARTIFACT_NAME=app.war
ARG JAR_FILE_PATH=build/libs/example-0.0.1-SNAPSHOTT.war

WORKDIR $APP_HOME

COPY --from=build $APP_HOME/$JAR_FILE_PATH $ARTIFACT_NAME

ENTRYPOINT ["java", "-jar", "app.war"]