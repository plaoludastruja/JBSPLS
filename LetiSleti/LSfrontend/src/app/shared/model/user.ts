export class User {
    id: string = '';
    firstName: string = '';
    lastName: string = '';
    email: string = '';
    password: string = '';

    public constructor(obj?: any) {
        if (obj) {
            this.id = obj.id;
            this.firstName = obj.firstName;
            this.lastName = obj.lastName;
            this.email = obj.email;
            this.password = obj.password;
        }
    }
}