<template>
    <q-scroll-area style="height: calc(100%);" :class="darkThemeSideBar[0] === 'class' ? darkThemeSideBar[1] : ''"
        :style="darkThemeSideBar[0] === 'style' ? darkThemeSideBar[1] : {}">
        <q-list>
            <template v-for="(childrenItem, index) in topMenuChildren" :key="index">
                <SideBarLeftItem :childrenItem="childrenItem" :initLevel="0" />
            </template>
        </q-list>
    </q-scroll-area>
</template>

<script setup>
import { toRefs } from 'vue';
import SideBarLeftItem from './SideBarLeftItem'
import useTheme from 'src/composables/useTheme';

const { darkThemeSideBar } = useTheme()
const props = defineProps({
    topMenuChildren: {
        type: Object || Array,
        required: false,
        default: () => {
            return null
        },
    },
})
const { topMenuChildren } = toRefs(props)
</script>
