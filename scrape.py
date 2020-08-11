import requests # Need to pip install
from bs4 import BeautifulSoup # Need to pip install
import shlex
import pandas as pd
'''
These packages must be installed
pip install pandas
pip install xlsxwriter
pip install xlrd
pip install openpyxl
'''

urls = ['https://thehearthaschambers.wordpress.com/category/songs/', 'https://thehearthaschambers.wordpress.com/category/songs/page/2/',
 'https://thehearthaschambers.wordpress.com/category/songs/page/3/', 'https://thehearthaschambers.wordpress.com/category/songs/page/4/']
songs = []
song_links = []
dates = []
title = []
slug = []

# Scrape all the song links from the website.
for url in urls:
    page_object = requests.get(url)
    soup = BeautifulSoup(page_object.text, 'html.parser')
    songs = soup.find_all(class_="entry-title")

    for song in songs:
        link = song.select('.entry-title a')
        title.append(link[0].text)
        song_links.append(link[0]['href'])
        new_link = link[0]['href'].replace("https://thehearthaschambers.wordpress.com/", "")
        # Split the remaining url link to get the date using rsplit. rsplit splits from the right side giving a specified
        # seperator. In this case it splits 2020/01/03/i-got-me/ into ['2020/01/03', 'i-got-me', '']
        dates.append(new_link.rsplit('/', 2)[0].replace("/", "-"))
        slug.append(new_link.rsplit('/', 2)[1])
        
#print(dates)
#print(song_links)

song_lyrics = []
for song_url in song_links:
    page_object = requests.get(song_url)
    soup = BeautifulSoup(page_object.text, 'html.parser')
    lyrics = soup.find_all(class_="entry-content")

    # Go through each of the lyrics and strip off the html tags
    for lyric in lyrics:
        dummy = lyric.text.rsplit('\n', 4)
        print(dummy[0])
        #print(lyric.text.rsplit('\n', 4)[0])
        song_lyrics.append(lyric.text.rsplit('\n', 4)[0])

#print(song_lyrics[0])
#print(dates[0])
#print(title[0])
#print(slug[0])

# Dataframe
#df = pd.DataFrame({'Title': title, 'Slug': slug, 'Body': song_lyrics, 'Dates': dates})

#writer = pd.ExcelWriter('demo.xlsx', engine='xlsxwriter')
#df.to_excel(writer, sheet_name='Sheet1', index=False)
#writer.save()

'''
id integer NOT NULL DEFAULT nextval('blog_id_seq'::regclass),
    author character varying(255) COLLATE pg_catalog."default",
    title character varying(255) COLLATE pg_catalog."default",
    slug character varying(255) COLLATE pg_catalog."default",
    body text COLLATE pg_catalog."default",
    created_date date NOT NULL DEFAULT CURRENT_DATE,
    published boolean NOT NULL,
'''