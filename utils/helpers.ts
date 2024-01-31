export const rootDomain = "http://localhost";

export const validateEmail = (email: string) => {
    const regex = new RegExp(/^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/);
    if (regex.test(email)){
        return true;
    }
    return false;
};