# frozen_string_literal: true

RSpec.describe Backoffice::ExchangeRatesController, type: :controller do
  describe '#collection' do
    context do
      before { subject.instance_variable_set :@collection, :collection }

      its(:collection) { should eq :collection }
    end

    context do
      let(:params) { { page: 14 } }

      before { expect(subject).to receive(:params).and_return(params) }

      before do
        #
        # ExchangeRate.order(date: :desc).page(14) -> collection
        #
        expect(ExchangeRate).to receive(:order).with(date: :desc) do
          double.tap do |a|
            expect(a).to receive(:page).with(14).and_return(:collection)
          end
        end
      end

      its(:collection) { should eq :collection }
    end
  end
end
