export const LocalStorage = {
    set : (keyName, newValue) => {
            window.localStorage.setItem(keyName, JSON.stringify(newValue));
    },
    get : (keyName) => {
        const value = window.localStorage.getItem(keyName);
        return value!==null ? JSON.parse(value) : {[keyName]: null};
    }
}