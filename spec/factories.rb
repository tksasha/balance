# frozen_string_literal: true

FactoryBot.define do
  factory :category do
    name { Faker::Commerce.product_name }

    currency { 'uah' }

    trait :usd do
      currency { 'usd' }
    end

    trait :uah do
      currency { 'uah' }
    end

    trait :eur do
      currency { 'eur' }
    end

    trait :visible do
      visible { true }
    end

    trait :invisible do
      visible { false }
    end

    trait :income do
      income { true }
    end

    trait :expense do
      income { false }
    end
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
      category factory: :category, income: true
    end

    trait :expense do
      category factory: :category, income: false
    end
  end

  factory :cash do
    name { SecureRandom.uuid }

    formula { '7 + 8' }

    trait :cash do
      supercategory { :cash }
    end

    trait :bonds do
      supercategory { :bonds }
    end

    trait :deposits do
      supercategory { :deposits }
    end
  end

  factory :tag do
    name { SecureRandom.uuid }

    category
  end
end
