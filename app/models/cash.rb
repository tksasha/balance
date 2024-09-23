# frozen_string_literal: true

# == Schema Information
#
# Table name: cashes
#
#  id            :integer          not null, primary key
#  currency      :integer          default("uah")
#  deleted_at    :time
#  favorite      :boolean          default(FALSE)
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

  enum :currency, CURRENCIES

  enum :supercategory, { cash: 1, bonds: 2, deposits: 3 }, default: :cash, scopes: false

  scope :favorite, -> { where(favorite: true) }

  has_paper_trail

  class << self
    def for_dashboard
      where(currency: CURRENCIES.keys)
        .group(:currency, :supercategory)
        .pluck('currency, supercategory, SUM(sum) AS sum')
        .group_by(&:first)
    end

    # TODO: spec me
    def ransackable_attributes(*)
      %w[currency deleted_at favorite formula id name sum supercategory]
    end
  end
end
