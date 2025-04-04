/**
 * Get the string of a month.
 * 
 * @param {number} month Month to consult.
 */
export function getMonthString(month: number): string {
    const months = ['Enero', 'Febrero', 'Marzo', 'Abril', 'Mayo', 'Junio', 'Julio', 'Agosto', 'Septiembre', 'Octubre', 'Noviembre', 'Diciembre'];
    return months[month];
}

/**
 * Get the month.
 */
export function getMonth(): string {
    let pastMonth: string;
    const currentDate = new Date();

    let month = currentDate.getMonth() + 1;
    if (month == 13) month = 1;

    let day = currentDate.getDay() - 1;
    if (day < 20) month = month + 1;

    if (month <= 9) pastMonth = `0${month}`;
    else pastMonth = `${month}`;
    return pastMonth;
}