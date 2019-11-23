FROM php:7.2-fpm

WORKDIR /app

# Install dependencies
RUN apt-get update && apt-get install -y \
    git \
    zip
RUN curl -sL https://getcomposer.org/installer | php -- --install-dir /usr/bin --filename composer
RUN docker-php-ext-install sockets

# Get application dependencies
COPY composer.json ./
RUN composer install -o

COPY main.php ./
