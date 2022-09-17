# frozen_string_literal: true

# == Schema Information
#
# Table name: categories
#
#  id            :integer          not null, primary key
#  currency      :integer          default("uah")
#  income        :boolean          default(FALSE)
#  name          :string
#  supercategory :integer          default("one"), not null
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

  enum :supercategory, { one: 1, two: 2, three: 3 }, default: :one, scopes: false

  scope :income, -> { where income: true }

  scope :expense, -> { where income: false }

  scope :visible, -> { where visible: true }

  class << self
    def supercategories
      defined_enums['supercategory']
        .each_with_object({}) do |i, res|
          name, = i

          res[I18n.t(name, scope: 'category.supercategory')] = name
        end
    end
  end
end
