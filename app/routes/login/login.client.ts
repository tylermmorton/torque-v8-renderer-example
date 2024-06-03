import {createSSRApp, h} from 'vue'

// @ts-ignore
import Login from "../../components/Login.vue";

export function mountApp(id: string, props: string) {
    console.log(`mounting app ${id} with props`, props)
    return createSSRApp(h(Login, JSON.parse(props))).mount(id, true)
}
