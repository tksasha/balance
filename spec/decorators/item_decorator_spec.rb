# frozen_string_literal: true

RSpec.describe ItemDecorator do
  subject { item.decorate }

  describe '#date' do
    context do
      let(:item) { stub_model Item, date: '2016-10-30' }

      its(:date) { should eq '30.10.2016' }
    end

    context do
      let(:item) { stub_model Item }

      its(:date) { should be_nil }
    end
  end
end
