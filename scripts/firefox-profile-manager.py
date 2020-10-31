#!/usr/bin/env python3

# # BASE_URL="https://download.mozilla.org/?product=firefox-${VERSION}&os=linux64&lang=en-US"
# # LATEST_STABLE_URL="https://download.mozilla.org/?product=firefox-latest&os=linux64&lang=en-US"

# function download() {
#    wget -O "firefox.tar.bz2" "https://download.mozilla.org/?product=firefox-latest&os=linux64&lang=en-US"
# }

# function create_profile() {
#    FIREFOX_EXECUTABLE=$1
#    PROFILE_PATH=$2

#    [ -z "$FIREFOX_EXECUTABLE" ] && return 1
#    [ -z "$PROFILE_PATH" ] && return 1

#    "$FIREFOX_EXECUTABLE" -profile "$PROFILE_PATH" &
#    # TODO get code from master branch of extract-browser-data.py
# }

# download "$1"

import shutil
import sys
import tarfile

import urllib3

URL_TEMPLATE = 'https://download.mozilla.org/?product={}&os=linux64&lang=en-US'
VERSION_LATEST = 'firefox-latest'
VERSION_BETA = 'firefox-beta-latest'
VERSION_ESR = 'firefox-esr-latest'
FILE_EXTENSION = '.tar.bz2'


def download_version(version, filepath):
   """Downloads Firefox version into filepath

   Returns success (bool) and response (str)
   """

   http = urllib3.PoolManager()
   r = http.request('GET', URL_TEMPLATE.format(version), preload_content=False)

   if r.status != 200:
      return False, r.reason

   with open(filepath, 'wb') as file:
      shutil.copyfileobj(r, file)

   r.release_conn()
   return True, r.reason


def extract(filepath, path):
   """Extracts the file to path, throws tarfile.ReadError if the file is not a bzip2 file"""
   with tarfile.open(filepath, "r:bz2") as file:
      file.extractall(path)


def create_profile(version, path):
   # downloads the version
   # runs it to create the profile (with xvfb if possible)
   # deletes the version
   #
   # do that for any specified version
   pass


extract('test.tar.bz2', '.')
# print('final',
#       download_into_file(VERSION_LATEST, 'firefox{}'.format(FILE_EXTENSION)))
