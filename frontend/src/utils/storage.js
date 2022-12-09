const storage = {
    set: (name, content, maxAge = null) => {
        if (!global.window || !name) {
            return
        }

        if (typeof content !== 'string') {
            content = JSON.stringify(content)
        }

        const storage = global.window.localStorage

        storage.setItem(name, content)
        if (maxAge && !isNaN(parseInt(maxAge))) {
            const timeout = parseInt(new Date().getTime() / 1000)
            storage.setItem(`${name}_EXPIRE`, timeout + maxAge)
        }
    },

    get: (name, defValue) => {
        defValue = defValue === undefined ? null : defValue
        if (!global.window || !name) {
            return defValue
        }
        const content = window.localStorage.getItem(name)

        if (!content) return defValue

        const _expire = window.localStorage.getItem(`${name}_EXPIRE`)
        if (_expire) {
            const now = parseInt(new Date().getTime() / 1000)
            if (now > _expire) {
                return defValue
            }
        }
        try {
            return JSON.parse(content)
        } catch (e) {
            return content
        }
    },

    rm: (name) => {
        if (!global.window || !name) {
            return
        }
        window.localStorage.removeItem(name)
        window.localStorage.removeItem(`${name}_EXPIRE`)
    },
    clear: () => {
        if (!global.window || !name) {
            return
        }
        window.localStorage.clear()
    }
}

export default storage