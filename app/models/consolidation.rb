# frozen_string_literal: true

class Consolidation < Item
  delegate :name, :income?, to: :category

  delegate :id, to: :category, prefix: true
end
