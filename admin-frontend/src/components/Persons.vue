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
                      aria-controls="addPersonModal"
                      @click="addPersonModal=true">Добавить</MDBBtn>
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
                  <tr class="fw-normal">
                    <th>
                        <div class="person-img-name-container">
                          <img
                              src="https://upload.wikimedia.org/wikipedia/commons/thumb/3/3e/Einstein_1921_by_F_Schmutzer_-_restoration.jpg/411px-Einstein_1921_by_F_Schmutzer_-_restoration.jpg"
                              alt="avatar 1" class="person-img">
                          <span>Альберт Эйнштейн</span>
                        </div>
                    </th>
                    <td class="align-middle text-center">
                      <h5 class="mb-0"><span class="badge bg-success">1</span></h5>
                    </td>
                    <td class="align-middle">
                      Эйнштейну принадлежит решающая роль в популяризации и введении в научный оборот новых физических...
                    </td>
                    <td class="align-middle">
                      <a href="#!" data-mdb-tooltip-init title="Done"><i
                          class="fas fa-edit fa-lg text-warning me-3"></i></a>
                      <a href="#!" data-mdb-tooltip-init title="Remove"><i
                          class="fas fa-trash-alt fa-lg text-danger"></i></a>
                    </td>
                  </tr>
                  <tr class="fw-normal">
                    <th>
                      <div class="person-img-name-container">
                        <img
                            src="https://upload.wikimedia.org/wikipedia/commons/thumb/7/79/Tesla_circa_1890.jpeg/411px-Tesla_circa_1890.jpeg"
                            alt="avatar 1" class="person-img">
                        <span>Никола Тесла</span>
                      </div>
                    </th>
                    <td class="align-middle text-center">
                      <h5 class="mb-0"><span class="badge bg-warning">2</span></h5>
                    </td>
                    <td class="align-middle">
                      Широко известен благодаря своему вкладу в создание устройств, работающих на переменном токе...
                    </td>
                    <td class="align-middle">
                      <a href="#!" data-mdb-tooltip-init title="Done"><i
                          class="fas fa-edit fa-lg text-warning me-3"></i></a>
                      <a href="#!" data-mdb-tooltip-init title="Remove"><i
                          class="fas fa-trash-alt fa-lg text-danger"></i></a>
                    </td>
                  </tr>
                  <tr class="fw-normal">
                    <th class="border-0">
                      <div class="person-img-name-container">
                        <img
                            src="https://upload.wikimedia.org/wikipedia/commons/thumb/e/ed/Elon_Musk_Royal_Society.jpg/400px-Elon_Musk_Royal_Society.jpg"
                            alt="avatar 1" class="person-img">
                        <span>Илон Маск</span>
                      </div>
                    </th>
                    <td class="border-0 align-middle"><h5 class="mb-0"><span class="badge bg-warning">2</span></h5></td>
                    <td class="border-0 align-middle">
                      В 2002 году Маск основал SpaceX, компанию по производству аэрокосмической...
                    </td>
                    <td class="border-0 align-middle">
                      <a href="#!" data-mdb-tooltip-init title="Done"><i
                          class="fas fa-edit fa-lg text-warning me-3"></i></a>
                      <a href="#!" data-mdb-tooltip-init title="Remove"><i
                          class="fas fa-trash-alt fa-lg text-danger"></i></a>
                    </td>
                  </tr>
                  </tbody>
                </table>


              </div>
            </div>

          </div>
        </div>
      </div>
    </section>
  </MDBContainer>


  <MDBModal
      id="addPersonModal"
      tabindex="-1"
      labelledby="addPersonModalLabel"
      size="lg"
      v-model="addPersonModal"
  >
    <MDBModalHeader>
      <MDBModalTitle id="addPersonModalLabel">Добавить личность</MDBModalTitle>
    </MDBModalHeader>
    <MDBModalBody>
      <MDBInput
          type="text"
          label="Личность"
          id="formName"
          v-model="formName"
          wrapperClass="mb-4"
      />

      <MDBInput
          type="email"
          label="Траектория"
          id="formEmail"
          v-model="formEmail"
          wrapperClass="mb-4"
      />

      <MDBTextarea
          label="Описание"
          id="formTextarea"
          v-model="formTextarea"
          wrapperClass="mb-4"
      />

      <MDBFile v-model="files" label="Добавить фото" />
    </MDBModalBody>
    <MDBModalFooter>
      <MDBBtn color="secondary" @click="addPersonModal = false">Закрыть</MDBBtn>
      <MDBBtn color="success">Сохранить</MDBBtn>
    </MDBModalFooter>
  </MDBModal>
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

<script setup lang="ts">
import {
  MDBContainer,
  MDBModal,
  MDBModalHeader,
  MDBModalTitle,
  MDBModalBody,
  MDBModalFooter,
  MDBBtn,
  MDBInput,
  MDBTextarea,
  MDBFile
} from "mdb-vue-ui-kit";

import { ref } from 'vue';
const addPersonModal = ref(false);

const formName = ref("");
const formEmail = ref("");
const formTextarea = ref("");
const files = ref([]);

defineProps<{ msg: string }>();
</script>
