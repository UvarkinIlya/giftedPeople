<script setup>
import {
  MDBBtn,
  MDBFile,
  MDBInput,
  MDBModal,
  MDBModalBody,
  MDBModalFooter,
  MDBModalHeader,
  MDBModalTitle,
  MDBTextarea
} from 'mdb-vue-ui-kit';
import {ref, watch} from 'vue';
import axios from 'axios';

const props = defineProps(["personData", "action"])
const emits = defineEmits(['close', 'update'])

console.log(props);
const isModalOpen = ref(false);

const personId = ref("");
const personName = ref("");
const personRoad = ref("");
const personDescription = ref("");
const personPhoto = ref();

if (props.action === "PUT"){
  watch(
      () => props.personData,
      (newVal) => {

        personId.value = newVal.id
        personName.value = newVal.name;
        personRoad.value = newVal.road;
        personDescription.value = newVal.description;
      }
  );
}


const savePerson = async () => {
  const person = {
    name: personName.value,
    road: personRoad.value,
    description: personDescription.value,
    file: personPhoto.value,
  }

  const formData = new FormData()
  formData.append("file[]", personPhoto.value)
  formData.append("name", personName.value)
  formData.append("road", personRoad.value)
  formData.append("description", personDescription.value)


  console.log(person);

  if (props.action === "PUT") {
    try {
      const response = await axios.put("http://" + window.location.hostname + ":8080/person/" + personId.value, formData)
    } catch (error) {
      console.log(error);
    }
  } else{
    try {
      const response = await axios.post("http://" + window.location.hostname + ":8080/person", formData);
    } catch (error) {
      console.log(error);
    }
  }

  emits('close')
  emits('update')

  personId.value = ""
  personName.value = ""
  personRoad.value = ""
  personDescription.value = ""
  personPhoto.value = ""
};
</script>

<template>
  <MDBModal
      id="addPersonModal"
      tabindex="-1"
      labelledby="addPersonModalLabel"
      size="lg"
      v-on:submit.prevent="savePerson"
      v-model="isModalOpen"
  >
    <form>
      <MDBModalHeader>
        <MDBModalTitle>{{props.action !== "PUT" ? "Добавить личность" : "Обновить личность"}}</MDBModalTitle>
      </MDBModalHeader>
      <MDBModalBody>
        <MDBInput
            type="text"
            label="Личность"
            wrapperClass="mb-4"
            v-model="personName"
        />

        <MDBInput
            type="text"
            label="Траектория"
            wrapperClass="mb-4"
            v-model="personRoad"
        />

        <MDBTextarea
            label="Описание"
            wrapperClass="mb-4"
            v-model="personDescription"
        />

        <MDBFile :label="props.action !== 'PUT' ? 'Добавить фото' : 'Обновить фото'" v-model="personPhoto"/>
      </MDBModalBody>
      <MDBModalFooter>
        <MDBBtn color="secondary" @click="$emit('close')">Закрыть</MDBBtn>
        <MDBBtn color="success" type="submit">Сохранить</MDBBtn>
      </MDBModalFooter>
    </form>
  </MDBModal>
</template>

<style scoped>

</style>