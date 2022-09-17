# frozen_string_literal: true

# == Schema Information
#
# Table name: cashes
#
#  id            :integer          not null, primary key
#  currency      :integer          default("uah")
#  deleted_at    :time
#  formula       :string
#  name          :string
#  sum           :decimal(10, 2)
#  supercategory :integer          default("cash"), not null
#
# Indexes
#
#  index_cashes_on_name_and_currency  (name,currency) UNIQUE
#

class Cash < ApplicationRecord
  include ActsAsHasFormula

  validates :name, :formula, :currency, presence: true

  validates :name, uniqueness: { case_sensitive: false, scope: :currency }

  enum currency: CURRENCIES

  enum :supercategory, { cash: 1, bonds: 2, deposits: 3 }, default: :cash, scopes: false

  has_paper_trail

  class << self
    def supercategories
      defined_enums['supercategory']
        .each_with_object({}) do |i, res|
          name, = i

          res[I18n.t(name, scope: 'cash.supercategory')] = name
        end
    end
  end
end
