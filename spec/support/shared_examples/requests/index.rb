# frozen_string_literal: true

RSpec.shared_examples 'index.js' do
  it { should render_template :index }

  it { expect(response).to have_http_status(:ok) }

  it { expect(response.content_type).to eq 'text/javascript; charset=utf-8' }
end
