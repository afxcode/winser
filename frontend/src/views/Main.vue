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
                                        <input class="w-9/12 fx-input" type="text" placeholder="Service Name" />
                                    </div>
                                    <div class="flex flex-row">
                                        <label class="w-3/12 fx-input-label "> Description </label>
                                        <input class="w-9/12 fx-input" type="text" placeholder="Description" />
                                    </div>
                                    <div class="flex flex-row">
                                        <label class="w-3/12 fx-input-label "> Startup Mode </label>
                                        <select class="w-9/12 fx-input" type="text" placeholder="Description">
                                            <option value="auto">Auto </option>
                                            <option value="manual">Manual</option>
                                        </select>
                                    </div>
                                    <div class="flex flex-row">
                                        <label class="w-3/12 fx-input-label "> Command </label>
                                        <input class="w-9/12 fx-input" type="text" placeholder="Command" />
                                    </div>
                                    <div class="flex flex-row">
                                        <label class="w-3/12 fx-input-label "> Working Dir </label>
                                        <input class="w-9/12 fx-input" type="text" placeholder="Working Dir" />
                                    </div>
                                </div>
                                <div class="flex flex-col w-1/5">
                                    <span class="fx-lbl-success"> RUNNING </span>
                                    <button class="fx-btn-success" type="button"> START </button>
                                    <button class="fx-btn-danger" type="button"> STOP </button>
                                    <button class="fx-btn-warning" type="button"> PAUSE </button>
                                </div>
                            </div>
                            <div class="flex flex-row w-full mt-2">
                                <button class="w-1/5 fx-btn-primary" :disabled="!state.btnFindEn" type="button">
                                    FIND</button>
                                <button class="w-1/5 fx-btn-primary" type="button"> EDIT </button>
                                <button class="w-1/5 fx-btn-success" type="button"> CREATE </button>
                                <button class="w-1/5 fx-btn-warning" type="button"> UPDATE </button>
                                <button class="w-1/5 fx-btn-danger" type="button"> REMOVE </button>
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

const state = reactive({
    btnFindEn: true,
})
</script>