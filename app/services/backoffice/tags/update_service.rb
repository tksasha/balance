# frozen_string_literal: true

module Backoffice
  module Tags
    class UpdateService < ApplicationService
      def initialize(params)
        @id = params.delete(:id)

        @params = params
      end

      def call
        return Success.new(tag) if tag.update(resource_params)

        Failure.new(tag)
      end

      private

      def tag
        @tag ||= Tag.find(@id)
      end

      def resource_params
        @params.require(:tag).permit(:name)
      end
    end
  end
end
