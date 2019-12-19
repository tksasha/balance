# frozen_string_literal: true

class ExchangeRate < ApplicationRecord
  validates :date, presence: true, uniqueness: { scope: %i[from to] }

  validates :from, presence: true

  validates :to, presence: true

  validates :rate, presence: true, numericality: { greater_than: 0 }

  enum from: CURRENCIES, _suffix: true

  enum to: CURRENCIES, _suffix: true
end
