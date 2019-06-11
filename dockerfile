FROM byuoitav/amd64-alpine
MAINTAINER Tony Petrangelo <tonypetrangelo@gmail.com>

ARG NAME
ENV name=${NAME}

COPY ${name}-bin ${name}-bin 
COPY version.txt version.txt

# add any required files/folders here

ENTRYPOINT ./${name}-bin
