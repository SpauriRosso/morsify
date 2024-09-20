#!/bin/bash

read -p "Enter your choice. Install(y) / Uninstall(n) / Cancel(c):- " choice

if [ "$choice" = "y" ]; then
  if [ -z "$(docker images -q morse-image:latest 2> /dev/null)" ]; then
    echo "morse-image does not exist"
    echo "Building image"
    docker build -t morse-image .
    echo "Starting Container"
    docker run -d --name morse -p 5827:5827 morse-image
    echo "Installation completed! Running on port 5827"
  else
    echo "morse-image:latest already exists"
    read -p "Do you want to launch the container (l) or uninstall the image (u)? " image_choice
    if [ "$image_choice" = "l" ]; then
      if [ "$(docker ps -q -f name=morse)" ]; then
        echo "Container 'morse' is already running"
      else
        echo "Starting Container"
        docker run -d --name morse -p 5827:5827 morse-image
        echo "Running on port 5827"
      fi
    elif [ "$image_choice" = "u" ]; then
      choice="n"
    else
      echo "Invalid choice. Exiting."
      exit 1
    fi
  fi
fi

if [ "$choice" = "n" ]; then
  echo "Uninstalling..."
  if [ "$(docker ps -q -f name=morse)" ]; then
    echo "Stopping and removing 'morse' container"
    docker stop morse
    docker rm morse
  fi
  if [ "$(docker images -q morse-image:latest)" ]; then
    echo "Removing 'morse-image' image"
    docker rmi morse-image:latest
  fi
  echo "Uninstallation completed"
elif [ "$choice" = "c" ]; then
  echo "Installation cancelled"
elif [ "$choice" != "y" ]; then
  echo "Wrong choice! Please enter y, n, or c."
fi
