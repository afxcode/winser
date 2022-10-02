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
                                        <input class="w-9/12 fx-input" type="text" v-model="state.name" placeholder="Service Name" />
                                    </div>
                                    <div class="flex flex-row">
                                        <label class="w-3/12 fx-input-label "> Display Name </label>
                                        <input class="w-9/12 fx-input" type="text" v-model="state.displayName" placeholder="Service Display Name" />
                                    </div>
                                    <div class="flex flex-row">
                                        <label class="w-3/12 fx-input-label "> Description </label>
                                        <textarea class="w-9/12 fx-input resize-none" type="text" rows="3" v-model="state.description" placeholder="Description" />
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
                                        <label class="w-3/12 fx-input-label "> Executable </label>
                                        <textarea class="w-8/12 fx-input resize-none mr-0" type="text" rows="2" v-model="state.executable" placeholder="Executable" />
                                        <button class="w-1/12 fx-btn-info h-8 ml-0 flex space-x-2 items-center" type="button" @click="selectExecutable"><FolderOpenIcon class="w-5 h-5"/></button>
                                    </div>
                                    <div class="flex flex-row">
                                        <label class="w-3/12 fx-input-label"> Argument </label>
                                        <textarea class="w-9/12 fx-input resize-none" type="file" rows="2" v-model="state.argument" placeholder="Argument" />
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
                                    <button class="fx-btn-success" type="button"> START </button>
                                    <button class="fx-btn-danger" type="button"> STOP </button>
                                    <button class="fx-btn-warning" type="button"> PAUSE </button>
                                </div>
                            </div>
                            <div class="flex flex-row w-full mt-2">
                                <button class="w-1/5 fx-btn-primary" type="button" :disabled="!state.btnFindEn" @click="find"> FIND </button>
                                <button class="w-1/5 fx-btn-primary" type="button"> EDIT </button>
                                <button class="w-1/5 fx-btn-success" type="button"> CREATE </button>
                                <button class="w-1/5 fx-btn-warning" type="button"> UPDATE </button>
                                <button class="w-1/5 fx-btn-danger" type="button"> REMOVE </button>
                            </div>
                            <div class="flex w-full mt-2">
                                <span class="fx-statusbar w-full"> {{state.statusbar}} </span>
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
import { Find, SelectExecutable } from '../../wailsjs/go/main/App'
import { FolderOpenIcon } from '@heroicons/vue/24/outline'

const state = reactive({
    name: "",
    displayName: "",
    description: "",
    executable: "",
    argument: "",
    startMode: "auto",

    status: "",

    btnFindEn: true,

    statusbar: "Ready",
})

function find() {
    Find(state.name).then(function (service) {
        state.displayName = service.DisplayName
        state.description = service.Description
        state.executable = service.Executable
        state.argument = service.Argument
        state.startMode = service.StartMode
        state.status = service.Status
        state.statusbar = "Find Success"
    }).catch(function (err) {
        state.statusbar = err
    })
}

function selectExecutable() {
    SelectExecutable(state.executable).then(function (selectedExecutable) {
        state.executable = selectedExecutable
    }).catch(function(err) {
        state.statusbar = err
    })
}
</script>