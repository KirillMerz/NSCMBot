package bot

const (
	REGISTER_HELP = "Чтобы зарегистрироваться, отправьте мне свои данные в формате:\n\n*\"Фамилия Имя Отчество 123456\"*, где \"123456\" - это номер паспорта.\n\nПри отсутствии отчества просто его не пишите.\nОбратите внимание на букву \"Ё\": заменять её на \"Е\" нельзя.\nДанные нужно вводить в одну строку, через пробел, без кавычек"

	MESSAGE_START = "Здесь вы можете узнать свои результаты экзаменов."

	MESSAGE_HELP = REGISTER_HELP + "\n\n/check - проверить результаты\n\n/unregister - удалить аккаунт\n\n/help - показать это сообщение\n\nПо всем вопросам обращайтесь к @KirillMerz"

	MESSAGE_REGISTER_SUCCESS = "Вы были успешно зарегистрированы\nЗапрашиваю результаты у NSCM..."

	MESSAGE_NOT_REGISTERED_ERROR = "Для этого вы должны зарегистрироваться :("

	MESSAGE_UNREGISTER_SUCCESS = "Ваши данные были успешно удалены"

	MESSAGE_DATABASE_ERROR = "В работе базы данных произошла ошибка. Пожалуйста, повторите попытку позже"

	MESSAGE_NSCM_IS_A_TEAPOT_ERROR = "Не удалось получить результаты. Повторите попытку позже или просто подождите, "

	MESSAGE_RESULTS_NOT_FOUND_ERROR = "Результаты по вашим данным не были найдены.\nПроверьте корректность введенных данных.\nЕсли вы ошиблись, удалите свой аккаунт (/unregister) и отправьте исправленные данные"

	MESSAGE_UNKNOWN_COMMAND_ERROR = "Я вас не понимаю...\nПоказать помощь - /help"
)
