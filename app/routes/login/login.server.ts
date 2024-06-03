import {createSSRApp,h} from 'vue'
import {renderToString} from '@vue/server-renderer'

// @ts-ignore
import Login from "../../components/Login.vue";

async function Render(props: string) {
    const app = createSSRApp(h(Login, JSON.parse(props)))

    return renderToString(app, {}).then((html) =>{
        return html
    });
}

(globalThis as any).Render = Render
