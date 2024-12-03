<template>
  <div class="personalities-page">
    <h1>Список личностей</h1>
    <button class="add-button" @click="openAddModal">Добавить новую личность</button>
    <div v-if="personalities.length" class="personalities-list">
      <div
          v-for="(personality, index) in personalities"
          :key="index"
          class="personality-card"
      >
        <img :src="personality.photo" alt="Фото личности" class="photo" />
        <div class="info">
          <p><strong>ФИО:</strong> {{ personality.name }}</p>
          <p><strong>Номер:</strong> {{ personality.phone }}</p>
          <p><strong>Описание:</strong> {{ personality.description }}</p>
        </div>
        <div class="actions">
          <button class="edit-button" @click="openEditModal(index)">Редактировать</button>
          <button class="delete-button" @click="deletePersonality(index)">Удалить</button>
        </div>
      </div>
    </div>
    <p v-else>Список пуст.</p>

    <!-- Модальное окно -->
    <div v-if="isModalOpen" class="modal-overlay">
      <div class="modal">
        <h2>{{ isEditing ? 'Редактировать личность' : 'Добавить новую личность' }}</h2>
        <form @submit.prevent="savePersonality">
          <label>
            Фото:
            <input type="text" v-model="currentPersonality.photo" placeholder="Ссылка на фото" />
          </label>
          <label>
            ФИО:
            <input type="text" v-model="currentPersonality.name" />
          </label>
          <label>
            Номер:
            <input type="text" v-model="currentPersonality.phone" />
          </label>
          <label>
            Описание:
            <textarea v-model="currentPersonality.description"></textarea>
          </label>
          <div class="modal-actions">
            <button class="save-button" type="submit">{{ isEditing ? 'Сохранить' : 'Добавить' }}</button>
            <button class="cancel-button" type="button" @click="closeModal">Отмена</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';

export default {
  name: 'PersonalitiesPage',
  setup() {
    const personalities = ref([]);
    const isModalOpen = ref(false);
    const isEditing = ref(false);
    const currentPersonality = ref({
      photo: '',
      name: '',
      phone: '',
      description: '',
    });
    const editingIndex = ref(null);

    const openAddModal = () => {
      resetCurrentPersonality();
      isEditing.value = false;
      isModalOpen.value = true;
    };

    const openEditModal = (index) => {
      currentPersonality.value = { ...personalities.value[index] };
      editingIndex.value = index;
      isEditing.value = true;
      isModalOpen.value = true;
    };

    const savePersonality = () => {
      if (isEditing.value) {
        personalities.value[editingIndex.value] = { ...currentPersonality.value };
      } else {
        personalities.value.push({ ...currentPersonality.value });
      }
      closeModal();
    };

    const deletePersonality = (index) => {
      personalities.value.splice(index, 1);
    };

    const closeModal = () => {
      isModalOpen.value = false;
      resetCurrentPersonality();
    };

    const resetCurrentPersonality = () => {
      currentPersonality.value = {
        photo: '',
        name: '',
        phone: '',
        description: '',
      };
    };

    return {
      personalities,
      isModalOpen,
      isEditing,
      currentPersonality,
      openAddModal,
      openEditModal,
      savePersonality,
      deletePersonality,
      closeModal,
    };
  },
};
</script>

<style scoped>
/* Общий стиль */
body {
  font-family: 'Arial', sans-serif;
  background-color: #f9f9f9;
  margin: 0;
  padding: 0;
}

.personalities-page {
  max-width: 900px;
  margin: 40px auto;
  padding: 20px;
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

/* Заголовок */
h1 {
  text-align: center;
  margin-bottom: 20px;
  color: #333;
}

/* Кнопка добавления */
.add-button {
  display: block;
  margin: 0 auto 20px auto;
  padding: 10px 20px;
  background: #4caf50;
  color: white;
  font-size: 16px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background 0.3s ease;
}

.add-button:hover {
  background: #45a049;
}

/* Список личностей */
.personalities-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 20px;
}

.personality-card {
  background: #fff;
  border: 1px solid #ddd;
  border-radius: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  padding: 15px;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.personality-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.photo {
  width: 100%;
  height: 150px;
  object-fit: cover;
  border-radius: 10px;
  margin-bottom: 10px;
}

.info {
  margin-bottom: 15px;
}

.actions button {
  padding: 8px 15px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 14px;
  transition: background 0.3s ease;
}

.edit-button {
  background: #2196f3;
  color: white;
}

.edit-button:hover {
  background: #1976d2;
}

.delete-button {
  background: #f44336;
  color: white;
}

.delete-button:hover {
  background: #d32f2f;
}

/* Модальное окно */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal {
  background: #fff;
  padding: 20px;
  border-radius: 10px;
  max-width: 500px;
  width: 100%;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.modal form label {
  display: block;
  margin-bottom: 10px;
  color: #333;
}

.modal form input,
.modal form textarea {
  width: 100%;
  padding: 10px;
  margin-bottom: 15px;
  border: 1px solid #ddd;
  border-radius: 5px;
}

.modal-actions {
  text-align: right;
}

.save-button {
  background: #4caf50;
  color: white;
}

.save-button:hover {
  background: #45a049;
}

.cancel-button {
  background: #f44336;
  color: white;
}

.cancel-button:hover {
  background: #d32f2f;
}
</style>
