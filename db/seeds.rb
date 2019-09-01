# frozen_string_literal: true

[
  ['Зарплата', true],
  ['%% в банці', true],
  ['Інше', true],
  ['Їжа', false],
  ['Проїзд', false],
  ["Дім. Сім'я", false],
  ['Непередбачуване', false],
  ['Автомобіль', false],
  ['Телефон', false],
  ['Квартира', false],
  ['Інтернет', false]
].each do |name, income|
  Category.create! name: name, income: income
end
