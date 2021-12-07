# frozen_string_literal: true

# == Schema Information
#
# Table name: categories
#
#  id       :integer          not null, primary key
#  currency :integer          default("uah")
#  income   :boolean          default(FALSE)
#  name     :string
#  visible  :boolean          default(TRUE)
#
# Indexes
#
#  index_categories_on_name_and_currency  (name,currency) UNIQUE
#
class Category < ApplicationRecord
  validates :name, presence: true, uniqueness: { case_sensitive: false, scope: :currency }

  validates :currency, presence: true

  enum currency: CURRENCIES

  scope :income, -> { where income: true }

  scope :expense, -> { where income: false }

  scope :visible, -> { where visible: true }
end
