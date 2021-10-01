import cron from 'node-cron';

import 'chromedriver';
import { Builder, By } from 'selenium-webdriver';
import chrome from 'selenium-webdriver/chrome';

import 'reflect-metadata';
import { createConnection } from 'typeorm';
import dayjs from 'dayjs';

import { Fortune } from './entity/fortune';

const zooKeyword = ['쥐띠', '소띠', '호랑이띠', '토끼띠', '용띠', '뱀띠', '말띠', '양띠', '원숭이띠', '닭띠', '개띠', '돼지띠'];
const constellKeyword = ['물병자리', '물고기자리', '양자리', '황소자리', '쌍둥이자리', '게자리', '사자자리', '처녀자리', '천칭자리', '전갈자리', '사수자리', '염소자리'];
const constellDates = ['1월 20일 ~ 2월 18일', '2월 19일 ~ 3월 20일', '3월 21일 ~ 4월 19일', '4월 20일 ~ 5월 20일', '5월 21일 ~ 6월 21일', '6월 22일 ~ 7월 22일', '7월 23일 ~ 8월 22일', '8월 23일 ~ 9월 23일', '9월 24일 ~ 10월 22일', '10월 23일 ~ 11월 22일', '11월 23일 ~ 12월 24일', '12월 25일 ~ 1월 19일'];

cron.schedule('0 45 23 * * *', () => {
  (async function crawler() {
    let data: Array<Fortune> = [];

    let driver = await new Builder().forBrowser('chrome').setChromeOptions(new chrome.Options().addArguments('headless')).build();

    await driver.manage().setTimeouts({ implicit: 10000, pageLoad: 30000, script: 30000 });

    try {
      for (let i = 0; i < zooKeyword.length; i++) {
        await driver.get(`https://search.naver.com/search.naver?where=nexearch&sm=top_hty&fbm=0&ie=utf8&query=${zooKeyword[i]}`);

        let fortuneContainer = await driver.findElement(By.id('yearFortune'));
        let selectFortunes = await driver.findElement(By.id('fortune_birthTab')).findElements(By.css('li > a'));

        await selectFortunes[1].click();

        try {
          await driver.wait(() => {
            return false;
          }, 1000);
        } catch (error) {

        }

        const fortunes = await fortuneContainer.findElements(By.css('p._cs_fortune_text'));
        console.log(zooKeyword[i]);

        data.push({
          name: zooKeyword[i],
          content: await fortunes[1].getText(),
          due_date: dayjs().add(1, 'day').format('YYYY-MM-DD'),
          created_at: dayjs().format('YYYY-MM-DD HH:mm:ss'),
          updated_at: dayjs().format('YYYY-MM-DD HH:mm:ss')

        });
      }

      for (let i = 0; i < constellKeyword.length; i++) {
        await driver.get(`https://search.naver.com/search.naver?where=nexearch&sm=top_hty&fbm=0&ie=utf8&query=${constellKeyword[i]}`);

        let fortuneContainer = await driver.findElement(By.id('yearFortune'));
        let selectFortunes = await fortuneContainer.findElement(By.css('ul.tab_wrap2._cs_fortune_tab')).findElements(By.css('li > a'));

        await selectFortunes[1].click();

        try {
          await driver.wait(() => {
            return false;
          }, 1000);
        } catch (error) {

        }

        const fortunes = await fortuneContainer.findElements(By.css('p._cs_fortune_text'));
        console.log(constellKeyword[i]);

        data.push({
          name: constellKeyword[i],
          content: `${constellDates[i]} ${await fortunes[1].getText()}`,
          due_date: dayjs().add(1, 'day').format('YYYY-MM-DD'),
          created_at: dayjs().format('YYYY-MM-DD HH:mm:ss'),
          updated_at: dayjs().format('YYYY-MM-DD HH:mm:ss')

        });
      }

    } finally {
      await driver.quit();
    }

    createConnection().then(async connection => {
      await connection.createQueryBuilder()
        .insert()
        .into(Fortune).values(data)
        .execute();

      console.log('Saving Data');
    }).catch(error => console.log(error));
  })();
});
