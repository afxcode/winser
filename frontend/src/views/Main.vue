<template>
  <main>
    <section class="absolute w-full h-full">
      <div class="container h-full">
        <div class="flex content-center justify-center h-full">
          <div class="relative flex flex-col min-w-0 break-words w-full bg-gray-300 border-0">
            <div class="rounded-t mb-0 px-2 py-2">
              <div class="flex flex-row w-full">
                <div class="flex flex-col w-4/5">
                  <div class="flex flex-row">
                    <label class="w-3/12 fx-input-label "> Service Name </label>
                    <input class="w-7/12 fx-input-sq rounded-l mr-0" type="text" v-model="state.name"
                      placeholder="Service Name" />
                    <button class="w-1/12 fx-btn-sq-info h-6 ml-0 mr-0 flex space-x-2 items-center" type="button"
                      :disabled="state.name == ''" @click="findService">
                      <MagnifyingGlassIcon class="w-6 h-4" />
                    </button>
                    <button class="w-1/12 fx-btn-sq-success rounded-r h-6 ml-0 flex space-x-2 items-center"
                      type="button" @click="selectListService">
                      <TableCellsIcon class="w-6 h-4" />
                    </button>
                  </div>
                  <div class="flex flex-row">
                    <label class="w-3/12 fx-input-label "> Display Name </label>
                    <input class="w-9/12 fx-input" type="text" v-model="state.displayName"
                      placeholder="Service Display Name" />
                  </div>
                  <div class="flex flex-row">
                    <label class="w-3/12 fx-input-label "> Description </label>
                    <textarea class="w-9/12 fx-input resize-none" type="text" rows="3" v-model="state.description"
                      placeholder="Description" />
                  </div>
                  <div class="flex flex-row">
                    <label class="w-3/12 fx-input-label "> Startup Mode </label>
                    <select class="w-9/12 fx-input" type="text" v-model="state.startMode">
                      <option disabled value="">Please select one</option>
                      <option value="auto">Auto </option>
                      <option value="manual">Manual</option>
                      <option value="disabled">Disabled</option>
                    </select>
                  </div>
                  <div class="flex flex-row">
                    <label class="w-3/12 fx-input-label"> Executable </label>
                    <textarea class="w-8/12 fx-input-sq rounded-l resize-none mr-0" type="text" rows="2"
                      v-model="state.executable" placeholder="Executable" :disabled="state.found" />
                    <button class="w-1/12 fx-btn-sq-info rounded-r h-8 ml-0 flex space-x-2 items-center" type="button"
                      :disabled="state.name == '' || state.found" @click="selectExecutable">
                      <FolderOpenIcon class="w-5 h-5" />
                    </button>
                  </div>
                  <div class="flex flex-row">
                    <label class="w-3/12 fx-input-label"> Argument </label>
                    <textarea class="w-9/12 fx-input resize-none" type="file" rows="2" v-model="state.argument"
                      placeholder="Argument" :disabled="state.found" />
                  </div>
                </div>
                <div class="flex flex-col w-1/5">
                  <span v-if="state.status == 'running'" class="fx-lbl-success"> RUNNING </span>
                  <span v-else-if="state.status == 'paused'" class="fx-lbl-info"> PAUSED </span>
                  <span v-else-if="state.status == 'stopped'" class="fx-lbl-info"> STOPPED </span>
                  <span v-else-if="state.status == 'pending_start'" class="fx-lbl-warning"> PENDING START </span>
                  <span v-else-if="state.status == 'pending_pause'" class="fx-lbl-warning"> PENDING PAUSE </span>
                  <span v-else-if="state.status == 'pending_stop'" class="fx-lbl-warning"> PENDING STOP </span>
                  <span v-else-if="state.status == 'pending_continue'" class="fx-lbl-warning"> PENDING CONTINUE </span>
                  <span v-else-if="state.status == 'unknown'" class="fx-lbl-danger"> UNKNOWN </span>
                  <span v-else class="fx-lbl-danger"> N/A </span>
                  <button class="fx-btn-success" type="button"
                    :disabled="state.status != 'stopped' && state.status != 'paused'"> START </button>
                  <button class="fx-btn-danger" type="button" :disabled="state.status != 'running'"> STOP </button>
                  <button class="fx-btn-warning" type="button" :disabled="state.status != 'running'"> PAUSE </button>
                </div>
              </div>
              <div class="flex flex-row w-full mt-2">
                <button class="w-3/12 fx-btn-secondary" type="button" @click="clear"> CLEAR </button>
                <button class="w-3/12 fx-btn-success" type="button" :disabled="!isFormValid()" @click="installService">
                  INSTALL </button>
                <button class="w-3/12 fx-btn-warning" type="button" :disabled="!isFormValid()" @click="updateService">
                  UPDATE </button>
                <button class="w-3/12 fx-btn-danger" type="button" :disabled="!isFormValid()" @click="removeService">
                  REMOVE </button>
              </div>
              <div class="flex w-full mt-2">
                <span class="w-full" :class="{
                    'fx-statusbar-primary': state.statusbarType == 'primary',
                    'fx-statusbar-success': state.statusbarType == 'success',
                    'fx-statusbar-warning': state.statusbarType == 'warning',
                    'fx-statusbar-danger': state.statusbarType == 'danger',
                }"> {{ state.statusbar }} </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </main>
</template>

<script lang="ts" setup>
import { reactive } from 'vue'
import { Find, SelectExecutable, Install, Remove, Update } from '../../wailsjs/go/main/App'
import { FolderOpenIcon, MagnifyingGlassIcon, TableCellsIcon } from '@heroicons/vue/24/outline'

const state = reactive({
  name: "",
  displayName: "",
  description: "",
  executable: "",
  argument: "",
  startMode: "",
  status: "",
  found: false,
  statusbar: "Ready",
  statusbarType: "primary"
})

function resetForm() {
  state.displayName = ""
  state.description = ""
  state.executable = ""
  state.argument = ""
  state.startMode = ""
  state.status = ""
  state.found = false
  state.statusbar = "Ready"
  state.statusbarType = "primary"
}

function isFormValid() {
  if (state.name == "") return false
  if (state.executable == "") return false
  if (state.startMode == "") return false
  return true
}

function clear() {
  state.name = ""
  resetForm()
}

function findService() {
  Find(state.name).then(function (service) {
    state.displayName = service.DisplayName
    state.description = service.Description
    state.executable = service.Executable
    state.argument = service.Argument
    state.startMode = service.StartMode
    state.status = service.Status
    state.found = true
    state.statusbar = "Find Success"
    state.statusbarType = "success"
  }).catch(function (err) {
    resetForm()
    state.statusbar = err
    state.statusbarType = "warning"
  })
}

function selectListService() {

}

function selectExecutable() {
  SelectExecutable(state.executable).then(function (selectedExecutable) {
    state.executable = selectedExecutable
    state.statusbar = "Service successfully selected"
    state.statusbarType = "success"
  }).catch(function (err) {
    state.statusbar = err
    state.statusbarType = "warning"
  })
}

function installService() {
  let service = {
    Name: state.name,
    DisplayName: state.displayName,
    Description: state.description,
    Executable: state.executable,
    Argument: state.argument,
    StartMode: state.startMode,
  }
  Install(service).then(function () {
    clear()
    state.statusbar = "Service successfully installed"
    state.statusbarType = "success"
  }).catch(function (err) {
    state.statusbar = err
    state.statusbarType = "danger"
  })
}

function updateService() {
  let service = {
    Name: state.name,
    DisplayName: state.displayName,
    Description: state.description,
    StartMode: state.startMode,
  }
  Update(service).then(function () {
    state.statusbar = "Service successfully updated"
    state.statusbarType = "success"
  }).catch(function (err) {
    state.statusbar = err
    state.statusbarType = "danger"
  })
}

function removeService() {
  Remove(state.name).then(function () {
    clear()
    state.statusbar = "Service successfully removed"
    state.statusbarType = "success"
  }).catch(function (err) {
    state.statusbar = err
    state.statusbarType = "danger"
  })
}
</script>