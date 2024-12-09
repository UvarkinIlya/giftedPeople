<script setup>
import {
  MDBContainer,
  MDBBtn,
} from 'mdb-vue-ui-kit';
import {nextTick, ref} from 'vue';
import ModalPerson from "./ModalPerson.vue";
import SinglePerson from "./SinglePerson.vue";

const isModelOpen = ref(false);
const persons = ref([]);

const update = async () => {
  const response = await fetch("http://localhost:8080/person");
  persons.value = await response.json();
  console.log(persons);
}

update()
</script>

<template>
  <MDBContainer fluid id="persons">
    <section class="min-vh-100 vw-100 gradient-custom-2">
      <div class="container py-5">
        <div class="row d-flex justify-content-center align-items-center h-100">
          <div class="col-md-12 col-xl-10">

            <div class="card mask-custom">
              <div class="card-body p-4 text-white">

                <div class="position-relative text-center pt-3 pb-2">
                  <h2 class="my-4">Одарённые личности</h2>
                  <MDBBtn
                      id="add-button"
                      class="btn btn-success position-absolute end-0 top-50 translate-middle-y"
                      color="success"
                      aria-controls="isModalOpen"
                      @click="isModelOpen=true">Добавить</MDBBtn>
                </div>
                <table class="table text-white mb-0">
                  <thead>
                  <tr>
                    <th scope="col">Личность</th>
                    <th class="align-middle text-center" scope="col">Траектория</th>
                    <th scope="col">Описание</th>
                    <th scope="col">Действия</th>
                  </tr>
                  </thead>
                  <tbody>
                  <SinglePerson
                      v-for="(person, index) in persons"
                      :key="person.id"
                      :name="person.name"
                      :img="person.img"
                      :road="person.road"
                      :description="person.desc"
                      :isLast="index === persons.length - 1"
                  />
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </MDBContainer>


  <ModalPerson v-model="isModelOpen" @update="update" @close="isModelOpen=false"/>
</template>

<style>
#persons {
  padding: 0;
  overflow-x: hidden;
}

#add-button {
  margin-top: 4px;
  padding-right: 2rem;
}

.person-img-name-container {
  display: flex;
  align-items: center;
  gap: 10px;
  padding-left: 0;
}

.person-img {
  width: 45px;
  height: auto;
  display: block;
}

th, td {
  text-align: center;
  vertical-align: middle;
}

.gradient-custom-2 {
  /* fallback for old browsers */
  background: #7e40f6;

  /* Chrome 10-25, Safari 5.1-6 */
  background: -webkit-linear-gradient(
      to right,
      rgba(126, 64, 246, 1),
      rgba(80, 139, 252, 1)
  );

  /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
  background: linear-gradient(
      to right,
      rgba(126, 64, 246, 1),
      rgba(80, 139, 252, 1)
  );
}

.mask-custom {
  background: rgba(24, 24, 16, 0.2);
  border-radius: 2em;
  backdrop-filter: blur(25px);
  border: 2px solid rgba(255, 255, 255, 0.05);
  background-clip: padding-box;
  box-shadow: 10px 10px 10px rgba(46, 54, 68, 0.03);
}
</style>
