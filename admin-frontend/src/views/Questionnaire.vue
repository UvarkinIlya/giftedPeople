<template>
  <div class="container">
    <h1 class="title">Тестирование</h1>

    <!-- Первый вопрос -->
    <div v-if="!answeredFirst" class="question">
      <!-- Текст вопроса -->
      <p class="question-text">{{ currentQuestion }}</p>

      <!-- Кнопки "Да" и "Нет" -->
      <div class="buttons" v-if="!showErrorFirst">
        <button class="yes-btn" @click="handleYesFirstClick">Да</button>
        <button class="no-btn" @click="handleNoFirstClick">Нет</button>
      </div>

      <!-- Сообщение об ошибке и кнопка "Начать тест заново" -->
      <div v-else>
        <p class="error-message">
          Данный тест не предназначен для выявления сферы одарённости и не может помочь в выборе стратегии её развития, если она не определена
        </p>
        <button class="restart-button" @click="restartTest">Начать тест заново</button>
      </div>
    </div>

    <!-- Второй вопрос -->
    <div v-if="answeredFirst && !answeredSecond" class="question">
      <p class="question-text">{{ nextQuestion }}</p>

      <div class="buttons">
        <button class="yes-btn" @click="handleYesSecondClick">Да</button>
        <button class="no-btn" @click="handleNoSecondClick">Нет</button>
      </div>
    </div>

    <!-- Третий вопрос -->
    <div v-if="answeredSecond && !answeredThird" class="question">
      <p class="question-text">{{ thirdQuestion }}</p>

      <div class="buttons">
        <button class="yes-btn" @click="handleYesThirdClick">Да</button>
        <button class="no-btn" @click="handleNoThirdClick">Нет</button>
      </div>
    </div>

    <!-- Четвёртый вопрос -->
    <div v-if="answeredThird && !answeredFourth" class="question">
      <p class="question-text">{{ fourthQuestion }}</p>

      <div class="buttons">
        <button class="yes-btn" @click="handleYesFourthClick">Да</button>
        <button class="no-btn" @click="handleNoFourthClick">Нет</button>
      </div>
    </div>

    <!-- Пятый вопрос -->
    <div v-if="answeredFourth && !answeredFifth && !logicalError" class="question">
      <p class="question-text">{{ fifthQuestion }}</p>
      <div class="buttons">
        <button class="yes-btn" @click="handleYesFifthClick">Да</button>
        <button class="no-btn" @click="handleNoFifthClick">Нет</button>
      </div>
    </div>

    <!-- Пятый вопрос, если логическая ошибка -->
    <div v-if="logicalError" class="question">
      <p class="error-message">
        Если у вас нет формального образования в вашей сфере одарённости, то ответ на второй вопрос не может быть «Да».
      </p>
      <button class="restart-button" @click="restartTest">Начать тест заново</button>
    </div>

    <!-- Шестой вопрос -->
    <div v-if="answeredFifth && !answeredSixth && !logicalError" class="question">
      <p class="question-text">{{ sixthQuestion }}</p>
      <div class="buttons vertical">
        <button class="yes-btn" @click="handleYesSixthClick">Моя сфера приложений усилий – традиционная</button>
        <button class="no-btn" @click="handleNoSixthClick">Я ищу новаторские пути, хочу создать совершенно новую отрасль/продукт</button>
      </div>
    </div>

    <!-- Седьмой вопрос -->
    <div v-if="answeredSixth && !answeredSeventh" class="question">
      <p class="question-text">{{ seventhQuestion }}</p>

      <div class="buttons vertical">
        <button class="yes-btn" @click="handleYesSeventhClick">Я – одиночка, у меня личный бренд</button>
        <button class="no-btn" @click="handleNoSeventhClick">Моя результативность и успех целиком и полностью зависит от того, есть ли у меня команда. Без команды (коллектива) я не достигну успеха и не смогу достичь самореализации</button>
      </div>
    </div>

    <!-- Сообщение о завершении теста -->
    <div v-if="answeredSeventh" class="question">
      <p class="cluster-description">
        {{ clusterDescription }}
      </p>

      <!-- Вывод данных о пользователе -->
      <div v-if="randomUser">
        <h3>{{ randomUser.name }}</h3>
        <p>{{ randomUser.desc }}</p>
<!--        <img :src="'http://' + window.location.hostname + ':8080/person-image/' + randomUser.img" alt="Фото пользователя" class="user-image">-->
        <button @click="restartTest" class="restart-button">Начать тест заново</button>
      </div>
    </div>



  </div>
</template>

<script>
import axios from 'axios';
export default {
  data() {
    return {
      answeredFirst: false,
      answeredSecond: false,
      answeredThird: false,
      answeredFourth: false,
      answeredFifth: false,
      answeredSixth: false,
      answeredSeventh: false,
      currentQuestion: 'Можете ли вы уверенно сказать, что именно является вашей «доминантой оригинальности» (сферой одарённости)?',
      nextQuestion: 'Можно ли сказать, что вы родились в обеспеченной семье?',
      thirdQuestion: 'Помогали ли ваши родители, или близкий родственный круг, в вашей сфере одарённости?',
      fourthQuestion: 'У вас есть формальное образование, которое соответствует вашей сфере одарённости?',
      fifthQuestion: 'Ваше первое формальное образование соответствует вашей сфере одарённости?',
      sixthQuestion: 'Моя сфера приложений усилий – традиционная / Я ищу новаторские пути, хочу создать совершенно новую отрасль/продукт',
      seventhQuestion: 'Я – одиночка, у меня личный бренд / Моя результативность и успех целиком и полностью зависит от того, есть ли у меня команда. Без команды (коллектива) я не достигну успеха и не смогу достичь самореализации',
      showErrorFirst: false,
      answerFirst: null,
      answerSecond: null,
      answerThird: null,
      answerFourth: null,
      answerFifth: null,
      answerSixth: null,
      answerSeventh: null,
      x1: null,
      x2: null,
      x3: null, // Переменная для расчета x3
      logicalError: false,
      clusterNumber: null, // Для отображения номера кластера
      clusterDescription: '', // Для описания кластера
      allUsers: [],
      randomUser: null,
      userImage: null,
    };
  },
  methods: {
    restartTest() {
      // Сброс всех переменных, связанных с ответами на вопросы
      this.answeredFirst = false;
      this.answeredSecond = false;
      this.answeredThird = false;
      this.answeredFourth = false;
      this.answeredFifth = false;
      this.answeredSixth = false;
      this.answeredSeventh = false;

      // Сброс всех ответов
      this.answerFirst = null;
      this.answerSecond = null;
      this.answerThird = null;
      this.answerFourth = null;
      this.answerFifth = null;
      this.answerSixth = null;
      this.answerSeventh = null;

      // Сброс всех вычислений
      this.x1 = null;
      this.x2 = null;
      this.x3 = null;

      // Сброс флага логической ошибки
      this.logicalError = false;

      // Сброс описания кластера и случайного пользователя
      this.clusterNumber = null;
      this.clusterDescription = '';
      this.randomUser = null;
      this.userImage = null;

      // Перезапуск теста — возвращаем начальный вопрос
      this.currentQuestion = 'Можете ли вы уверенно сказать, что именно является вашей «доминантой оригинальности» (сферой одарённости)?';
      this.showErrorFirst = false;

      // Получаем случайного пользователя для нового теста
      this.getRandomUserFromCluster(this.clusterNumber);
    },
    async getAllUsers() {
      try {
        const response = await axios.get('http://' + window.location.hostname + ':8080/person');
        this.allUsers = response.data;
      } catch (error) {
        console.error("Ошибка при получении данных пользователей", error);
      }
    },

    // Метод для получения случайного пользователя из кластера
    async getRandomUserFromCluster(clusterNumber) {
      try {
        const response = await axios.get('http://' + window.location.hostname + `:8080/person/by-road-id/${clusterNumber}`);
        this.randomUser = response.data;

        // После получения пользователя, загружаем его изображение
        // if (this.randomUser && this.randomUser.img) {
        //   // this.getUserImage(this.randomUser.img); // Передаем imgId для загрузки фото
        // }
      } catch (error) {
        console.error("Ошибка при получении случайного пользователя", error);
      }
    },

    // Метод для получения фото пользователя
    // async getUserImage(imgId) {
    //   try {
    //     const response = await axios.get('http://' + window.location.hostname + `:8080/person-image/${imgId}`);
    //     this.userImage = response.data;
    //   } catch (error) {
    //     console.error("Ошибка при получении фото пользователя", error);
    //   }
    // },
    mounted() {
      // Получаем всех пользователей при загрузке компонента
      this.getAllUsers();

      // Выбираем случайного пользователя из второго кластера
      this.getRandomUserFromCluster(this.clusterNumber);
    },
    handleYesFirstClick() {
      this.answerFirst = 'yes';
      this.answeredFirst = true;
      this.showErrorFirst = false;
    },
    handleNoFirstClick() {
      this.answerFirst = 'no';
      this.showErrorFirst = true;
    },
    handleYesSecondClick() {
      this.answerSecond = 'yes';
      this.answeredSecond = true;
      this.calculateX1();
    },
    handleNoSecondClick() {
      this.answerSecond = 'no';
      this.answeredSecond = true;
      this.calculateX1();
    },
    handleYesThirdClick() {
      this.answerThird = 'yes';
      this.answeredThird = true;
      this.calculateX1();
    },
    handleNoThirdClick() {
      this.answerThird = 'no';
      this.answeredThird = true;
      this.calculateX1();
    },
    handleYesFourthClick() {
      this.answerFourth = 'yes';
      this.answeredFourth = true;
      this.calculateX2();
    },
    handleNoFourthClick() {
      this.answerFourth = 'no';
      this.answeredFourth = true;
      this.calculateX2();
    },
    handleYesFifthClick() {
      this.answerFifth = 'yes';
      this.answeredFifth = true;
      this.calculateX2();
    },
    handleNoFifthClick() {
      this.answerFifth = 'no';
      this.answeredFifth = true;
      this.calculateX2();
    },
    handleYesSixthClick() {
      this.answerSixth = 'yes';  // Традиционная сфера приложений усилий
      this.answeredSixth = true;
      this.calculateX3();
    },
    handleNoSixthClick() {
      this.answerSixth = 'no';  // Новаторские пути
      this.answeredSixth = true;
      this.calculateX3();
    },
    handleYesSeventhClick() {
      this.answerSeventh = 'yes';  // Личный бренд (Одиночка)
      this.answeredSeventh = true;
      this.calculateX3();
      this.clusterDescription = this.getClusterDescription();
      this.mounted()
    },
    handleNoSeventhClick() {
      this.answerSeventh = 'no';  // Зависимость от команды
      this.answeredSeventh = true;
      this.calculateX3();
      this.clusterDescription = this.getClusterDescription();
      this.mounted()
    },
    calculateX1() {
      if (this.answerSecond === 'yes' && this.answerThird === 'yes') {
        this.x1 = 4;
      } else if (this.answerSecond === 'no' && this.answerThird === 'no') {
        this.x1 = 1;
      } else if (this.answerSecond === 'yes' && this.answerThird === 'no') {
        this.x1 = 3;
      } else if (this.answerSecond === 'no' && this.answerThird === 'yes') {
        this.x1 = 2;
      }
      console.log(`x1 = ${this.x1}`);
      this.calculateCluster();
    },
    calculateX2() {
      if (this.answerFourth === 'yes' && this.answerFifth === 'yes') {
        this.x2 = 3;
      } else if (this.answerFourth === 'yes' && this.answerFifth === 'no') {
        this.x2 = 2;
      } else if (this.answerFourth === 'no' && this.answerFifth === 'no') {
        this.x2 = 1;
      } else if (this.answerFourth === 'no' && this.answerFifth === 'yes') {
        this.logicalError = true;
        this.x2 = null;
      }
      if (!this.logicalError) {
        console.log(`x2 = ${this.x2}`);
        this.calculateCluster();
      }
    },
    calculateX3() {
      if (this.answerSixth === 'yes' && this.answerSeventh === 'yes') {
        this.x3 = 1;  // Традиционная сфера приложений усилий + Личный бренд (Одиночка)
      } else if (this.answerSixth === 'yes' && this.answerSeventh === 'no') {
        this.x3 = 2;  // Традиционная сфера приложений усилий + Зависимость от команды
      } else if (this.answerSixth === 'no' && this.answerSeventh === 'yes') {
        this.x3 = 3;  // Новаторские пути + Личный бренд (Одиночка)
      } else if (this.answerSixth === 'no' && this.answerSeventh === 'no') {
        this.x3 = 4;  // Новаторские пути + Зависимость от команды
      }

      console.log(`x3 = ${this.x3}`);
      this.calculateCluster();  // После расчета x3, запускаем расчет кластера
    },

    calculateCluster() {
      if (this.x1 !== null && this.x2 !== null && this.x3 !== null) {
        const clusters = [
          { name: 'Голубой', x1_mean: 3.11, x2_mean: 3, x3_mean: 2.44, x1_std: 0.33, x2_std: 0.01, x3_std: 0.88 },
          { name: 'Оранжевый', x1_mean: 1.54, x2_mean: 2.77, x3_mean: 2.15, x1_std: 0.52, x2_std: 0.44, x3_std: 0.55 },
          { name: 'Серый', x1_mean: 3.75, x2_mean: 1.7, x3_mean: 3.65, x1_std: 0.44, x2_std: 0.47, x3_std: 0.49 },
          { name: 'Желтый', x1_mean: 3.67, x2_mean: 3, x3_mean: 4, x1_std: 0.49, x2_std: 0.01, x3_std: 0.01 },
          { name: 'Синий', x1_mean: 1.15, x2_mean: 3, x3_mean: 3.7, x1_std: 0.37, x2_std: 0.01, x3_std: 0.47 },
          { name: 'Зеленый', x1_mean: 1.27, x2_mean: 1.8, x3_mean: 3.67, x1_std: 0.46, x2_std: 0.41, x3_std: 0.62 },
          { name: 'Темно-синий', x1_mean: 4, x2_mean: 2.88, x3_mean: 2.5, x1_std: 0.01, x2_std: 0.35, x3_std: 0.53 }
        ];

        let minDistance = Infinity;
        let clusterNumber = -1;

        console.log(`Входные значения: x1=${this.x1}, x2=${this.x2}, x3=${this.x3}`);
        for (let i = 0; i < clusters.length; i++) {
          const cluster = clusters[i];

          // Вычисляем сумму отклонений
          const sumDifferences =
              Math.abs(this.x1 - cluster.x1_mean) +
              Math.abs(this.x2 - cluster.x2_mean) +
              Math.abs(this.x3 - cluster.x3_mean);

          console.log(
              `Кластер ${cluster.name}: отклонение = ${sumDifferences.toFixed(2)}`
          );

          if (sumDifferences < minDistance) {
            minDistance = sumDifferences;
            clusterNumber = i + 1; // Нумерация кластеров с 1
          }
        }


        this.clusterNumber = clusterNumber;
        console.log(`Итоговый кластер: ${clusterNumber}`);
      }
    },
    getClusterDescription() {
      switch (this.clusterNumber) {
        case 1:
          return  "Первый кластер («Инноваторы традиций») характеризуется хорошими стартовыми условиями, точным самоопределением. Большинство представителей этой группы достигли успеха в естественно-научной сфере. Как правило, на жизненном пути одарённые из этого кластера отличались большой самостоятельностью (например, А. Эйнштейн или Э. Шрёдингер).";
        case 2:
          return "Во втором кластере («Одиссеи науки») − преимущество за гуманитариями. Представители этого кластера имели непростые стартовые условия, но отличались быстрым самоопределением на этапе обучения. Однако сфера приложения их усилий не отличается высокой степенью инновационности и высокой научной коллаборацией. Здесь мы найдём много «одиночек», самостоятельно сформировавшихся как профессионалы в науке. Средняя скорость успеха – почти 36 лет, что существенно выше среднего значения по всей выборке. Например, представителем этого кластера является биография Ф. Ницше.";
        case 3:
          this.clusterDescription =returnbreak;
        case 4:
          return "Попавших в четвёртый кластер можно характеризовать как «круглых отличников», «Моцарты науки». У них хороший старт и точное самоопределение, сфера применения их усилий отличается высокой инновационностью. Не случайно средний возраст успеха в этом кластере 31,5 года. Большинство представителей этого кластера работали в сфере точных и естественных наук. Например, в этот кластер попадают биографии Н. Бора и М. Борна.";
        case 5:
          return "В пятый кластер («Стойкие экспериментаторы») попали ученые, чьи биографии отличались самыми сложными стартовыми условиями, которые им приходилось восполнять «удачей» и трудолюбием на втором и третьем этапе. Здесь много экспериментаторов – например, Д. Менделеев, А. Макаренко. Профессиональный путь таких людей требует времени и большого трудолюбия, но достижения имеют инновационное значение для отрасли.";
        case 6:
          return "В шестом кластере сосредоточены «загадочные таланты»: их путь представляется самым неопределённым с точки зрения условий для успеха. Для них характерны сложные стартовые условия и низкая скорость самоопределения в период обучения и становления. Вторая учеба, или второстепенная работа, «мешала» сосредоточиться на доминанте оригинальности. При этом их включение в научную коммуникацию ничем не уступает «экспериментаторскому» пути. Пример: Л. Пастер, К. Д. Ушинский.";
        case 7:
          return "Представители седьмого кластера («Прорывные мастера») отличаются самыми лучшими стартовыми условиями и самым коротким временем до первого успеха – всего 28,75 лет. Это связано с рано проявленной гениальностью и высокой оригинальностью работ. В этом кластере есть как физики, так и гуманитарии. Пример: М. Планк, Л. Ландау, Ж.-П. Сартр.";
        default:
          return "Неверное значение для кластера.";
      }

      return this.clusterDescription;
    }

  },
};
</script>

<style scoped>
.container {
  max-width: 600px;
  margin: 0 auto;
  text-align: center;
  padding: 20px;
}

.title {
  font-size: 24px;
  margin-bottom: 20px;
}

.question {
  margin-bottom: 20px;
}

.question-text {
  font-size: 18px;
  margin-bottom: 15px;
}

.buttons {
  display: flex;
  justify-content: center;
  gap: 20px;
}

.yes-btn, .no-btn {
  padding: 10px 20px;
  font-size: 16px;
  cursor: pointer;
}

.yes-btn {
  background-color: #4CAF50;
  color: white;
  border: none;
}

.no-btn {
  background-color: #f44336;
  color: white;
  border: none;
}

.vertical .yes-btn, .vertical .no-btn {
  display: block;
  width: 100%;
  margin-bottom: 10px;
}

.error-message {
  color: red;
  margin-top: 10px;
}

.user-image {
  max-width: 400px; /* Уменьшает максимальную ширину изображения */
  max-height: 400px; /* Уменьшает максимальную высоту изображения */
  display: block;
  margin: 0 auto; /* Центрирует изображение по горизонтали */
}
.restart-button {
  margin-top: 20px;
  padding: 10px 20px;
  font-size: 16px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
.restart-button:hover {
  background-color: #45a049;
}
</style>

