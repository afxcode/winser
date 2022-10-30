<template>
  <main>
    <section class="absolute w-full h-full">
      <div class="container h-full">
        <div class="flex content-center justify-center h-full">
          <div class="relative flex flex-col min-w-0 break-words w-full bg-gray-300 border-0">
            <div v-if="state.screen == 'main'" class="rounded-t mb-0 px-2 py-2">
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
            <div v-if="state.screen == 'select'" class="rounded-t mb-0 px-2 py-2">
              <div class="flex flex-row w-full">
                <table class="w-full text-xs text-left text-gray-500 dark:text-gray-400">
                  <thead class="flex flex-col text-xs text-gray-700 bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                    <tr class="flex w-full">
                      <th scope="col" class="px-4 py-2 w-5/12"> Name </th>
                      <th scope="col" class="px-4 py-2 w-4/12"> Display </th>
                      <th scope="col" class="px-4 py-2 w-2/12"> Status </th>
                      <th scope="col" class="px-1 pt-1  w-1/12">
                        <button class="flex mx-0 mb-0 px-1 w-full fx-btn-secondary" type="button"
                          @click="state.screen = 'main'">
                          <ArrowUturnLeftIcon class="w-full h-4" />
                        </button>
                      </th>
                    </tr>
                  </thead>
                  <tbody class="text-xs overflow-y-auto bg-grey-light flex flex-col" style="height: 85vh">
                    <tr v-for="(srv, idx) in state.listServices" :key="srv.name" class="flex w-full bg-white">
                      <td class="px-4 py-1 w-5/12"> {{srv.name}} </td>
                      <td class="px-4 py-1 w-4/12"> {{srv.display}} </td>
                      <td class="px-2 py-0 w-2/12">
                        <span v-if="srv.status == 'running'" class="fx-lbl-success w-full mt-1 mb-1"> RUNNING </span>
                        <span v-if="srv.status == 'paused'" class="fx-lbl-info w-full mt-1 mb-1"> PAUSED </span>
                        <span v-if="srv.status == 'stopped'" class="fx-lbl-info w-full mt-1 mb-1"> STOPPED </span>
                        <span v-if="srv.status == 'pending_start'" class="fx-lbl-warning w-full mt-1 mb-1"> P START
                        </span>
                        <span v-if="srv.status == 'pending_pause'" class="fx-lbl-warning w-full mt-1 mb-1"> P PAUSE
                        </span>
                        <span v-if="srv.status == 'pending_stop'" class="fx-lbl-warning w-full mt-1 mb-1"> P STOP
                        </span>
                        <span v-if="srv.status == 'pending_continue'" class="fx-lbl-warning w-full mt-1 mb-1"> P CONT
                        </span>
                        <span v-if="srv.status == 'unknown'" class="fx-lbl-danger w-full mt-1 mb-1"> UNKNOWN </span>
                      </td>
                      <td class="px-2 py-0 w-1/12 mr-2">
                        <button class="flex w-12/12 h-6 mt-1 fx-btn-info" type="button" @click="viewService(srv.name)">
                          <EyeIcon class="w-4 h-4" />
                        </button>
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
  </main>
</template>

<script lang="ts" setup>
import { reactive } from 'vue'
import { GetServices, Find, SelectExecutable, Install, Remove, Update } from '../../wailsjs/go/main/App'
import { FolderOpenIcon, MagnifyingGlassIcon, TableCellsIcon, EyeIcon, ArrowUturnLeftIcon } from '@heroicons/vue/24/outline'

type ListService = {
  name: string
  display: string
  status: string
}
const listServices: ListService[] = []

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
  statusbarType: "primary",

  listServices: listServices,

  screen: "main",
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
  GetServices().then(function (services) {
    let srvs: ListService[] = []
    for (let i = 0; i < services.length; i++) {
      const srv = services[i]
      srvs.push({ name: srv.Name, display: srv.DisplayName, status: srv.Status })
    }
    state.listServices = srvs
    state.screen = 'select'
  }).catch(function (err) {
    resetForm()
    state.statusbar = err
    state.statusbarType = "warning"
  })
}

function viewService(name: string) {
  state.name = name
  state.screen = 'main'
  if (name == "") return
  findService()
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