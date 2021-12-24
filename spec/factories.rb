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
