![banner](https://github.com/11notes/defaults/blob/main/static/img/banner.png?raw=true)

# SYNCTHING
![size](https://img.shields.io/docker/image-size/11notes/syncthing/2.0.9?color=0eb305)![5px](https://github.com/11notes/defaults/blob/main/static/img/transparent5x2px.png?raw=true)![version](https://img.shields.io/docker/v/11notes/syncthing/2.0.9?color=eb7a09)![5px](https://github.com/11notes/defaults/blob/main/static/img/transparent5x2px.png?raw=true)![pulls](https://img.shields.io/docker/pulls/11notes/syncthing?color=2b75d6)![5px](https://github.com/11notes/defaults/blob/main/static/img/transparent5x2px.png?raw=true)[<img src="https://img.shields.io/github/issues/11notes/docker-SYNCTHING?color=7842f5">](https://github.com/11notes/docker-SYNCTHING/issues)![5px](https://github.com/11notes/defaults/blob/main/static/img/transparent5x2px.png?raw=true)![swiss_made](https://img.shields.io/badge/Swiss_Made-FFFFFF?labelColor=FF0000&logo=data:image/svg%2bxml;base64,PHN2ZyB2ZXJzaW9uPSIxIiB3aWR0aD0iNTEyIiBoZWlnaHQ9IjUxMiIgdmlld0JveD0iMCAwIDMyIDMyIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPgogIDxyZWN0IHdpZHRoPSIzMiIgaGVpZ2h0PSIzMiIgZmlsbD0idHJhbnNwYXJlbnQiLz4KICA8cGF0aCBkPSJtMTMgNmg2djdoN3Y2aC03djdoLTZ2LTdoLTd2LTZoN3oiIGZpbGw9IiNmZmYiLz4KPC9zdmc+)

Run syncthing rootless and distroless.

# INTRODUCTION 📢

Syncthing is a continuous file synchronization program. It synchronizes files between two or more computers.

![DASHBOARD](https://github.com/11notes/docker-syncthing/blob/master/img/Dashboard.png?raw=true)

# SYNOPSIS 📖
**What can I do with this?** This image will run syncthing [rootless](https://github.com/11notes/RTFM/blob/main/linux/container/image/rootless.md) and [distroless](https://github.com/11notes/RTFM/blob/main/linux/container/image/distroless.md), for maximum security and performance. If no configuration is found this image will automatically generate a new one with the environment variables used. This image will also by default disable telemetry.

# UNIQUE VALUE PROPOSITION 💶
**Why should I run this image and not the other image(s) that already exist?** Good question! Because ...

> [!IMPORTANT]
>* ... this image runs [rootless](https://github.com/11notes/RTFM/blob/main/linux/container/image/rootless.md) as 1000:1000
>* ... this image has no shell since it is [distroless](https://github.com/11notes/RTFM/blob/main/linux/container/image/distroless.md)
>* ... this image is auto updated to the latest version via CI/CD
>* ... this image has a health check
>* ... this image runs read-only
>* ... this image is automatically scanned for CVEs before and after publishing
>* ... this image is created via a secure and pinned CI/CD process
>* ... this image is very small
>* ... this image has a custom init process for more comfort

If you value security, simplicity and optimizations to the extreme, then this image might be for you.

# COMPARISON 🏁
Below you find a comparison between this image and the most used or original one.

| **image** | **size on disk** | **init default as** | **[distroless](https://github.com/11notes/RTFM/blob/main/linux/container/image/distroless.md)** | supported architectures
| ---: | ---: | :---: | :---: | :---: |
| linuxserver/syncthing | 58MB | 0:0 | ❌ | amd64, arm64 |

# VOLUMES 📁
* **/syncthing/etc** - Directory of the configuration file
* **/syncthing/var** - Directory of database and index data
* **/syncthing/share** - Directory of the default share (can be used as mount point for multiple shares)

# COMPOSE ✂️
```yaml
name: "syncthing"
services:
  server:
    image: "11notes/syncthing:2.0.9"
    read_only: true
    environment:
      TZ: "Europe/Zurich"
      SYNCTHING_PASSWORD: "${SYNCTHING_PASSWORD}"
      SYNCTHING_API_KEY: "${SYNCTHING_API_KEY}"
    volumes:
      - "syncthing.etc:/syncthing/etc"
      - "syncthing.var:/syncthing/var"
      - "syncthing.share:/syncthing/share"
    ports:
      - "3000:3000/tcp"
      - "22000:22000/tcp"
      - "22000:22000/udp"
      - "21027:21027/udp"
    networks:
      frontend:
    restart: "always"

volumes:
  syncthing.etc:
  syncthing.var:
  syncthing.share:

networks:
  frontend:
```
To find out how you can change the default UID/GID of this container image, consult the [how-to.changeUIDGID](https://github.com/11notes/RTFM/blob/main/linux/container/image/11notes/how-to.changeUIDGID.md#change-uidgid-the-correct-way) section of my [RTFM](https://github.com/11notes/RTFM)

# DEFAULT SETTINGS 🗃️
| Parameter | Value | Description |
| --- | --- | --- |
| `user` | docker | user name |
| `uid` | 1000 | [user identifier](https://en.wikipedia.org/wiki/User_identifier) |
| `gid` | 1000 | [group identifier](https://en.wikipedia.org/wiki/Group_identifier) |
| `home` | /syncthing | home directory of user docker |
| `username` | admin | username |

# ENVIRONMENT 📝
| Parameter | Value | Default |
| --- | --- | --- |
| `TZ` | [Time Zone](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) | |
| `DEBUG` | Will activate debug option for container image and app (if available) | |
| `SYNCTHING_PASSWORD` | Password for the admin user, **must** be set! | |
| `SYNCTHING_API_KEY` | API key, **must** be set and 32 characters long! | |

# MAIN TAGS 🏷️
These are the main tags for the image. There is also a tag for each commit and its shorthand sha256 value.

* [2.0.9](https://hub.docker.com/r/11notes/syncthing/tags?name=2.0.9)

### There is no latest tag, what am I supposed to do about updates?
It is of my opinion that the ```:latest``` tag is dangerous. Many times, I’ve introduced **breaking** changes to my images. This would have messed up everything for some people. If you don’t want to change the tag to the latest [semver](https://semver.org/), simply use the short versions of [semver](https://semver.org/). Instead of using ```:2.0.9``` you can use ```:2``` or ```:2.0```. Since on each new version these tags are updated to the latest version of the software, using them is identical to using ```:latest``` but at least fixed to a major or minor version.

If you still insist on having the bleeding edge release of this app, simply use the ```:rolling``` tag, but be warned! You will get the latest version of the app instantly, regardless of breaking changes or security issues or what so ever. You do this at your own risk!

# REGISTRIES ☁️
```
docker pull 11notes/syncthing:2.0.9
docker pull ghcr.io/11notes/syncthing:2.0.9
docker pull quay.io/11notes/syncthing:2.0.9
```

# SOURCE 💾
* [11notes/syncthing](https://github.com/11notes/docker-SYNCTHING)

# PARENT IMAGE 🏛️
> [!IMPORTANT]
>This image is not based on another image but uses [scratch](https://hub.docker.com/_/scratch) as the starting layer.
>The image consists of the following distroless layers that were added:
>* [11notes/distroless](https://github.com/11notes/docker-distroless/blob/master/arch.dockerfile) - contains users, timezones and Root CA certificates
>* [11notes/distroless:localhealth](https://github.com/11notes/docker-distroless/blob/master/localhealth.dockerfile) - app to execute HTTP requests only on 127.0.0.1

# BUILT WITH 🧰
* [syncthing](https://github.com/syncthing/syncthing)

# GENERAL TIPS 📌
> [!TIP]
>* Use a reverse proxy like Traefik, Nginx, HAproxy to terminate TLS and to protect your endpoints
>* Use Let’s Encrypt DNS-01 challenge to obtain valid SSL certificates for your services

# ElevenNotes™️
This image is provided to you at your own risk. Always make backups before updating an image to a different version. Check the [releases](https://github.com/11notes/docker-syncthing/releases) for breaking changes. If you have any problems with using this image simply raise an [issue](https://github.com/11notes/docker-syncthing/issues), thanks. If you have a question or inputs please create a new [discussion](https://github.com/11notes/docker-syncthing/discussions) instead of an issue. You can find all my other repositories on [github](https://github.com/11notes?tab=repositories).

*created 14.09.2025, 07:26:12 (CET)*