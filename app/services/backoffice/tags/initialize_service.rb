# frozen_string_literal: true

module Backoffice
  module Tags
    class InitializeService < ApplicationService
      def initialize(category)
        @category = category
      end

      def call
        Success.new(tag)
      end

      private

      def tag
        @category.tags.new
      end
    end
  end
end
