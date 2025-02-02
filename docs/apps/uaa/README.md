# Deploy UAA as a Kf App

These instructions will walk you through deploying [Cloud Foundry UAA][uaa] as a Kf app.

## Deploy

1. Download and extract the UAA source:

    ```sh
    wget https://github.com/cloudfoundry/uaa/archive/4.35.0.zip
    unzip 4.35.0.zip
    cd uaa-*
    ```

1. Add a default servlet mapping for Tomcat (this is necessary as the
   Tomcat Cloud Native Buildpack does not include a global web.xml with this
   mapping.):

    ```sh
    cat >uaa/src/main/webapp/WEB-INF/tomcat-web.xml <<EOF
    <?xml version="1.0" encoding="ISO-8859-1"?>
    <web-app xmlns="http://java.sun.com/xml/ns/javaee" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
      xsi:schemaLocation="http://java.sun.com/xml/ns/javaee http://java.sun.com/xml/ns/javaee/web-app_3_0.xsd" version="3.0"
      metadata-complete="true">
      <servlet>
            <servlet-name>default</servlet-name>
            <servlet-class>
              org.apache.catalina.servlets.DefaultServlet
            </servlet-class>
            <init-param>
                <param-name>debug</param-name>
                <param-value>0</param-value>
            </init-param>
            <init-param>
                <param-name>listings</param-name>
                <param-value>false</param-value>
            </init-param>
            <load-on-startup>1</load-on-startup>
        </servlet>
    </web-app>
    EOF
    ```

1. Create a sample deployment manifest. This snippet contains a minimally functional UAA configuration based on the manifest generated by running `gradlew manifests` in the UAA source repo. You should replace the `UAA_CONFIG_YAML` value with your own (see https://github.com/cloudfoundry/uaa/issues/238 for details on creating your own UAA config from scratch):

    ```sh
    cat >manifest.yml <<EOF
    ---
    applications:
    - name: uaa
      env:
        BP_JAVA_VERSION: 8.*
        BP_BUILT_ARTIFACT: uaa/build/libs/cloudfoundry-identity-uaa-*.war
        UAA_URL: http://localhost:8080
        LOGIN_URL: http://localhost:8080
        UAA_CONFIG_YAML: |
          issuer:
            uri: http://localhost:8080
          encryption:
            active_key_label: CHANGE-THIS-KEY
            encryption_keys:
            - label: CHANGE-THIS-KEY
              passphrase: CHANGEME
          login:
            serviceProviderKey: |
              -----BEGIN RSA PRIVATE KEY-----
              MIICXQIBAAKBgQDHtC5gUXxBKpEqZTLkNvFwNGnNIkggNOwOQVNbpO0WVHIivig5
              L39WqS9u0hnA+O7MCA/KlrAR4bXaeVVhwfUPYBKIpaaTWFQR5cTR1UFZJL/OF9vA
              fpOwznoD66DDCnQVpbCjtDYWX+x6imxn8HCYxhMol6ZnTbSsFW6VZjFMjQIDAQAB
              AoGAVOj2Yvuigi6wJD99AO2fgF64sYCm/BKkX3dFEw0vxTPIh58kiRP554Xt5ges
              7ZCqL9QpqrChUikO4kJ+nB8Uq2AvaZHbpCEUmbip06IlgdA440o0r0CPo1mgNxGu
              lhiWRN43Lruzfh9qKPhleg2dvyFGQxy5Gk6KW/t8IS4x4r0CQQD/dceBA+Ndj3Xp
              ubHfxqNz4GTOxndc/AXAowPGpge2zpgIc7f50t8OHhG6XhsfJ0wyQEEvodDhZPYX
              kKBnXNHzAkEAyCA76vAwuxqAd3MObhiebniAU3SnPf2u4fdL1EOm92dyFs1JxyyL
              gu/DsjPjx6tRtn4YAalxCzmAMXFSb1qHfwJBAM3qx3z0gGKbUEWtPHcP7BNsrnWK
              vw6By7VC8bk/ffpaP2yYspS66Le9fzbFwoDzMVVUO/dELVZyBnhqSRHoXQcCQQCe
              A2WL8S5o7Vn19rC0GVgu3ZJlUrwiZEVLQdlrticFPXaFrn3Md82ICww3jmURaKHS
              N+l4lnMda79eSp3OMmq9AkA0p79BvYsLshUJJnvbk76pCjR28PK4dV1gSDUEqQMB
              qy45ptdwJLqLJCeNoR0JUcDNIRhOCuOPND7pcMtX6hI/
              -----END RSA PRIVATE KEY-----
            serviceProviderKeyPassword: password
            serviceProviderCertificate: |
              -----BEGIN CERTIFICATE-----
              MIIDSTCCArKgAwIBAgIBADANBgkqhkiG9w0BAQQFADB8MQswCQYDVQQGEwJhdzEO
              MAwGA1UECBMFYXJ1YmExDjAMBgNVBAoTBWFydWJhMQ4wDAYDVQQHEwVhcnViYTEO
              MAwGA1UECxMFYXJ1YmExDjAMBgNVBAMTBWFydWJhMR0wGwYJKoZIhvcNAQkBFg5h
              cnViYUBhcnViYS5hcjAeFw0xNTExMjAyMjI2MjdaFw0xNjExMTkyMjI2MjdaMHwx
              CzAJBgNVBAYTAmF3MQ4wDAYDVQQIEwVhcnViYTEOMAwGA1UEChMFYXJ1YmExDjAM
              BgNVBAcTBWFydWJhMQ4wDAYDVQQLEwVhcnViYTEOMAwGA1UEAxMFYXJ1YmExHTAb
              BgkqhkiG9w0BCQEWDmFydWJhQGFydWJhLmFyMIGfMA0GCSqGSIb3DQEBAQUAA4GN
              ADCBiQKBgQDHtC5gUXxBKpEqZTLkNvFwNGnNIkggNOwOQVNbpO0WVHIivig5L39W
              qS9u0hnA+O7MCA/KlrAR4bXaeVVhwfUPYBKIpaaTWFQR5cTR1UFZJL/OF9vAfpOw
              znoD66DDCnQVpbCjtDYWX+x6imxn8HCYxhMol6ZnTbSsFW6VZjFMjQIDAQABo4Ha
              MIHXMB0GA1UdDgQWBBTx0lDzjH/iOBnOSQaSEWQLx1syGDCBpwYDVR0jBIGfMIGc
              gBTx0lDzjH/iOBnOSQaSEWQLx1syGKGBgKR+MHwxCzAJBgNVBAYTAmF3MQ4wDAYD
              VQQIEwVhcnViYTEOMAwGA1UEChMFYXJ1YmExDjAMBgNVBAcTBWFydWJhMQ4wDAYD
              VQQLEwVhcnViYTEOMAwGA1UEAxMFYXJ1YmExHTAbBgkqhkiG9w0BCQEWDmFydWJh
              QGFydWJhLmFyggEAMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEEBQADgYEAYvBJ
              0HOZbbHClXmGUjGs+GS+xC1FO/am2suCSYqNB9dyMXfOWiJ1+TLJk+o/YZt8vuxC
              KdcZYgl4l/L6PxJ982SRhc83ZW2dkAZI4M0/Ud3oePe84k8jm3A7EvH5wi5hvCkK
              RpuRBwn3Ei+jCRouxTbzKPsuCVB+1sNyxMTXzf0=
              -----END CERTIFICATE-----
          LOGIN_SECRET: loginsecret
          jwt:
            policy:
              activeKeyId: key-id-1
              keys:
                key-id-1:
                  signingKey: |
                    -----BEGIN RSA PRIVATE KEY-----
                    MIICXQIBAAKBgQDHtC5gUXxBKpEqZTLkNvFwNGnNIkggNOwOQVNbpO0WVHIivig5
                    L39WqS9u0hnA+O7MCA/KlrAR4bXaeVVhwfUPYBKIpaaTWFQR5cTR1UFZJL/OF9vA
                    fpOwznoD66DDCnQVpbCjtDYWX+x6imxn8HCYxhMol6ZnTbSsFW6VZjFMjQIDAQAB
                    AoGAVOj2Yvuigi6wJD99AO2fgF64sYCm/BKkX3dFEw0vxTPIh58kiRP554Xt5ges
                    7ZCqL9QpqrChUikO4kJ+nB8Uq2AvaZHbpCEUmbip06IlgdA440o0r0CPo1mgNxGu
                    lhiWRN43Lruzfh9qKPhleg2dvyFGQxy5Gk6KW/t8IS4x4r0CQQD/dceBA+Ndj3Xp
                    ubHfxqNz4GTOxndc/AXAowPGpge2zpgIc7f50t8OHhG6XhsfJ0wyQEEvodDhZPYX
                    kKBnXNHzAkEAyCA76vAwuxqAd3MObhiebniAU3SnPf2u4fdL1EOm92dyFs1JxyyL
                    gu/DsjPjx6tRtn4YAalxCzmAMXFSb1qHfwJBAM3qx3z0gGKbUEWtPHcP7BNsrnWK
                    vw6By7VC8bk/ffpaP2yYspS66Le9fzbFwoDzMVVUO/dELVZyBnhqSRHoXQcCQQCe
                    A2WL8S5o7Vn19rC0GVgu3ZJlUrwiZEVLQdlrticFPXaFrn3Md82ICww3jmURaKHS
                    N+l4lnMda79eSp3OMmq9AkA0p79BvYsLshUJJnvbk76pCjR28PK4dV1gSDUEqQMB
                    qy45ptdwJLqLJCeNoR0JUcDNIRhOCuOPND7pcMtX6hI/
                    -----END RSA PRIVATE KEY-----
    EOF
    ```

1. Deploy (this assumes you've already `kf target`'d a space):

    ```sh
    kf push uaa
    ```

1. Use the proxy feature to access the deployed app:

    ```sh
    kf proxy uaa
    ```

1. In a separate terminal, curl the endpoint via the proxy:

    ```sh
    curl 'http://localhost:8080/info' -i -H 'Accept: application/json'
    ```

    The response should indicate a successful deployment:

    ```sh
    HTTP/1.1 200 OK
    Cache-Control: no-store
    Content-Language: en-US
    Content-Type: application/json;charset=UTF-8
    Date: Tue, 16 Jul 2019 22:14:07 GMT
    Server: istio-envoy
    X-Envoy-Upstream-Service-Time: 26
    Content-Length: 383

    {"app":{"version":"4.35.0"},"links":{"uaa":"http://localhost:8080","passwd":"/forgot_password","login":"http://localhost:8080","register":"/create_account"},"zone_name":"uaa","entityID":"cloudfoundry-saml-login","commit_id":"git-metadata-not-found","idpDefinitions":{},"prompts":{"username":["text","Email"],"password":["password","Password"]},"timestamp":"2019-07-16T14:56:04-0700"
    ```

## Destroy

1. Delete the app:

    ```sh
    kf delete uaa
    ```

[uaa]: https://github.com/cloudfoundry/uaa
