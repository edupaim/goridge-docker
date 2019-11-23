FROM php:7.2-fpm

WORKDIR /app

RUN apt-get update && apt-get install -y \
    git \
    zip
RUN docker-php-ext-install sockets
RUN curl -sL https://getcomposer.org/installer | php -- --install-dir /usr/bin --filename composer

COPY composer.json ./
RUN composer install -o
COPY main.php ./

RUN echo '#!/bin/bash\nphp main.php' > /usr/bin/run && \
    chmod +x /usr/bin/run
