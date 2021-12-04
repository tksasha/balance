# frozen_string_literal: true

FactoryBot.define do
  factory :category do
    name { Faker::Commerce.department }

    currency { 'uah' }
  end

  factory :item do
    date { Date.new 2019, 11, 13 }

    category

    formula { '2 + 3' }
  end

  factory :cash do
    name { Faker::Lorem.word }

    formula { '7 + 8' }
  end
end
