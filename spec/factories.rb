# frozen_string_literal: true

FactoryBot.define do
  factory :category do
    name { SecureRandom.uuid }

    currency { 'uah' }
  end

  factory :item do
    date { Time.zone.today }

    category

    formula { '2 + 3' }

    trait :uah do
      currency { 'uah' }
    end

    trait :usd do
      currency { 'usd' }
    end

    trait :eur do
      currency { 'eur' }
    end

    trait :income do
      association :category, income: true
    end

    trait :expense do
      association :category, income: false
    end
  end

  factory :cash do
    name { SecureRandom.uuid }

    formula { '7 + 8' }
  end

  factory :tag do
    name { SecureRandom.uuid }

    category
  end
end
