# frozen_string_literal: true

module Backoffice
  module Tags
    class GetCollectionService < ApplicationService
      def initialize(category)
        @category = category
      end

      def call
        @category.tags.order(:name)
      end
    end
  end
end
