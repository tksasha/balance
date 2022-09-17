# frozen_string_literal: true

# == Schema Information
#
# Table name: categories
#
#  id            :integer          not null, primary key
#  currency      :integer          default("uah")
#  income        :boolean          default(FALSE)
#  name          :string
#  supercategory :integer          default("first"), not null
#  visible       :boolean          default(TRUE)
#
# Indexes
#
#  index_categories_on_name_and_currency  (name,currency) UNIQUE
#

class Category < ApplicationRecord
  validates :name, presence: true, uniqueness: { case_sensitive: false, scope: :currency }

  validates :currency, presence: true

  enum currency: CURRENCIES

  enum :supercategory, { first: 1, second: 2, third: 3 }, default: :first, scopes: false

  scope :income, -> { where income: true }

  scope :expense, -> { where income: false }

  scope :visible, -> { where visible: true }
end
